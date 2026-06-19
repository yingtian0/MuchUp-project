# 0002. Use slog API with zerolog backend for structured logging

## Status

Accepted

## Date

2026-06-11

## Context


`app` はリアルタイムチャットの中核を担い、以下のような処理を扱う。

* HTTP API
* WebSocket
* gRPC
* Redis
* PostgreSQL / GORM
* ユーザー認証
* チャットメッセージ処理
* ルーム単位のイベント配信

このため、`app` のログは単なる文字列ではなく、検索・集計・監視・障害調査に使える構造化ログとして出力する必要がある。

特に以下のようなフィールドをログに含めたい。

| 対象           | ログに含める主な情報                                            |
| ------------ | ----------------------------------------------------- |
| HTTP request | method, path, status, latency_ms, request_id, user_id |
| WebSocket    | user_id, room_id, remote_addr, event                  |
| chat message | message_id, room_id, sender_id                        |
| Redis stream | stream, event_type, message_id                        |
| gRPC         | service, method, duration_ms, error                   |
| DB           | operation, table, latency_ms, error                   |

現在、`app` では `github.com/rs/zerolog` を導入済みであり、`app/pkg/logger` に `zerolog.Logger` をラップした独自 logger が存在する。

一方で、現在の logger API には以下のような課題がある。

* `Infof`, `Errorf`, `Warnf` など printf 系 API が残っている
* `Info(msg string, args ...any)` も内部では `Msgf` を使っている
* key-value 形式の構造化ログとして扱いにくい
* `WithContext(ctx)` が context の値を logger に反映していない
* 一部の handler で標準 `log.Printf` が混在している
* `authz` 側は `app` とログ方針が統一されていない

MuchUp は WebSocket や Redis Stream など高頻度ログが発生しうる構成であり、今後のパフォーマンス要求も考慮する必要がある。
そのため、標準 API としての扱いやすさ、将来の差し替えやすさ、実行時性能のバランスを取る必要がある。

## Decision

MuchUp-project では、ログ方針を以下のように定める。

1. `app` は構造化 JSON ログを標準とする
2. `app` のログ出力 backend は `zerolog` を使う
3. handler / usecase / service などのアプリケーションコードでは、`slog` に寄せた API を使う
4. `zerolog` への直接依存は logger package や infrastructure 層に閉じ込める
5. その他の小さな Go サービスでは、原則として `slog` API を使う
6. `authz` も `log.Printf` / `log.Fatal` ではなく、`slog` または共通 logger package に寄せる

最終的な構成は以下を目指す。

```text
Application code / handlers / usecases
    ↓
slog-compatible logging API
    ↓
zerolog backend
    ↓
JSON logs to stdout
    ↓
Docker / Kubernetes / Cloud Logging / Loki / Datadog
```

`app` 内部では、今後のパフォーマンス要求を考慮し、ログ出力 backend として `zerolog` を維持する。

ただし、handler や usecase が `zerolog` の fluent API に直接依存する形にはしない。

アプリケーション側のログ記述は、以下のような key-value 形式に統一する。

```go
logger.Info(
    "user connected",
    "user_id", userID,
    "room_id", roomID,
)
```

エラーログも同様に、文字列へ埋め込まず属性として渡す。

```go
logger.Error(
    "failed to publish message",
    "error", err,
    "user_id", userID,
    "room_id", roomID,
    "message_id", messageID,
)
```

中期的には、以下のように `slog.Logger` をアプリケーション側APIとし、`zerolog` をhandler/backendとして利用する構成へ寄せる。

```go
zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
handler := zerolog.NewSlogHandler(zl)
logger := slog.New(handler)

logger.Info(
    "user connected",
    "user_id", userID,
    "room_id", roomID,
)
```

ただし、利用する `zerolog` の slog handler 実装は、導入時点で保守状況・互換性・出力形式を確認してから採用する。

## Alternatives considered

### Use zerolog directly everywhere

`zerolog` を handler / usecase / infrastructure すべてで直接使う案。

メリット:

* 高速な JSON ログ出力を直接利用できる
* 既存の `app/pkg/logger` と親和性が高い
* 追加の抽象化が少ない

デメリット:

* アプリケーションコードが `zerolog` に強く依存する
* 将来 `slog.JSONHandler` や OpenTelemetry 対応 handler に差し替えにくい
* handler / usecase に `log.Info().Str(...).Msg(...)` のような backend 固有 API が広がる
* 小さいサービスとのログ API 統一が難しくなる

このため、`zerolog` の直接利用は logger package や infrastructure 層に限定する。

### Use slog.JSONHandler only

標準ライブラリの `slog.JSONHandler` に全面移行する案。

メリット:

* 標準ライブラリだけで完結する
* API が安定している
* 小さいサービスでは導入しやすい
* handler / usecase の記述がシンプルになる

デメリット:

* 既存の `zerolog` 導入を捨てる理由が薄い
* `app` のような高頻度ログが発生しうるサービスでは、将来的な性能面の余地を狭める
* 既存 logger package の移行コストが発生する

このため、`authz` など小さいサービスでは選択肢に残すが、`app` の主要 backend としては採用しない。


### Use slog API with zerolog backend

アプリケーション側は `slog` に寄せ、出力 backend として `zerolog` を使う案。

メリット:

* handler / usecase は標準的な key-value logging API を使える
* backend として `zerolog` の高速 JSON 出力を利用できる
* 将来 `slog.JSONHandler` や OpenTelemetry 対応 handler に差し替えやすい
* logger の実装詳細をアプリケーション層から隠せる
* `app` とその他サービスでログ記述スタイルを統一しやすい

デメリット:

* `slog` と `zerolog` を接続する handler の選定・検証が必要
* 既存 logger package の API 見直しが必要
* 移行期間中は既存 API と新 API が一時的に併存する可能性がある

この案が、標準 API・性能・将来の拡張性のバランスが最も良いため採用する。

## Consequences

### Positive

* `app` のログが JSON 構造化ログとして安定して出力される
* HTTP / WebSocket / gRPC / Redis / DB のログに共通フィールドを付与しやすくなる
* `request_id`, `user_id`, `room_id`, `trace_id`, `span_id` などを context から自動付与できる
* 障害調査時に、ユーザー単位・ルーム単位・リクエスト単位でログを追いやすくなる
* handler / usecase が `zerolog` 固有 API に依存しなくなる
* 将来的に OpenTelemetry や外部ログ基盤へ接続しやすくなる
* `authz` など他サービスでも `slog` ベースの統一方針を適用しやすくなる

### Negative

* 既存の `Infof`, `Errorf`, `Warnf` 利用箇所を段階的に置き換える必要がある
* `app/pkg/logger` の interface と実装を見直す必要がある
* `zerolog` 用 slog handler の導入・検証が必要になる
* 一時的に旧 logger API と新 logger API が混在する可能性がある
* ログフィールド名の命名規則を決めないと、構造化ログの品質がばらつく

### Required follow-up work

1. `app/pkg/logger` の API を key-value 形式へ寄せる

```go
type Logger interface {
    Debug(msg string, attrs ...any)
    Info(msg string, attrs ...any)
    Warn(msg string, attrs ...any)
    Error(msg string, attrs ...any)
    Fatal(msg string, attrs ...any)

    With(attrs ...any) Logger
    WithError(err error) Logger
    WithContext(ctx context.Context) Logger
}
```

2. `Infof`, `Errorf`, `Warnf`, `Fatalf` の新規利用を禁止する

既存コードは段階的に移行する。

悪い例:

```go
logger.Infof("user connected user_id=%s room_id=%s", userID, roomID)
```

良い例:

```go
logger.Info(
    "user connected",
    "user_id", userID,
    "room_id", roomID,
)
```

3. 標準 `log.Printf` / `log.Fatal` を排除する

WebSocket handler などに残っている標準 `log` を、共通 logger に置き換える。

例:

```go
logger.Error(
    "websocket read error",
    "error", err,
    "user_id", userID,
    "room_id", roomID,
)
```

4. `WithContext(ctx)` を実装する

context から以下の値を抽出してログに自動付与できるようにする。

* `request_id`
* `user_id`
* `room_id`
* `trace_id`
* `span_id`

5. HTTP request logging を構造化する

request metrics / access log は以下のような属性を持つ。

```go
logger.Info(
    "request completed",
    "method", r.Method,
    "path", r.URL.Path,
    "status", status,
    "latency_ms", latency.Milliseconds(),
    "request_id", requestID,
    "user_id", userID,
)
```

6. ログフィールド名の命名規則を定義する

原則として snake_case を使う。

例:

* `request_id`
* `user_id`
* `room_id`
* `message_id`
* `latency_ms`
* `remote_addr`
* `event_type`
* `service`
* `method`
* `status`
* `error`

7. `authz` のログ方針を統一する

`authz` は比較的小さいサービスであるため、以下のどちらかを採用する。

* `slog` のみを使う
* `app` と同じ共通 logger package を使う

ただし、標準 `log` の直接利用は避ける。

## Notes

このADRの主眼は、`slog` と `zerolog` のどちらか一方を選ぶことではない。

MuchUp-project における判断は以下である。

* `app` は構造化 JSON ログを必須とする
* `app` の内部 backend は、今後のパフォーマンス要求を考慮して `zerolog` を使う
* handler / usecase / service などのアプリケーションコードは、`slog` に寄せた key-value logging API を使う
* その他のサービスでは、標準 API として `slog` を優先する
* `zerolog` の詳細は logger package または infrastructure 層に閉じ込める

最終的なログ出力の流れは以下とする。

```text
Application code
    ↓
slog-compatible API
    ↓
zerolog backend
    ↓
JSON logs to stdout
    ↓
Docker / Kubernetes / Cloud Logging / Loki / Datadog
```
