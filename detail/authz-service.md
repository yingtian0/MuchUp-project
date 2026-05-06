# Auth Service (authz-service)

## 役割

リバースプロキシ（Envoy）での認証ゲートに加え、
アプリケーション用のログイン・サインアップ API を提供する。
JWT 発行と最小限のユーザー認証を担う。

## 提供 API

- `POST /auth/login`
  - 目的: 既存ユーザーのログイン
  - リクエスト: `email`, `password`
  - レスポンス: `token`, `userId`, `username`

- `POST /auth/signup`
  - 目的: 新規ユーザー登録
  - リクエスト: `username`, `email`, `password`
  - レスポンス: `token`, `userId`, `username`

- `GET /healthz`
  - 目的: ヘルスチェック
  - レスポンス: `200 OK`

## 認証・トークン

- JWT は HS256 で署名
- 署名鍵は環境変数 `JWT_SECRET` で上書き可能
- 有効期限は `JWT_TTL` で指定可能（例: `24h`）

## データ管理

- ユーザー情報はインメモリ保持
- 永続化は将来の拡張対象

## エラー

- `400`: リクエスト不正
- `401`: 認証失敗
- `409`: ユーザー重複
- `500`: 内部エラー
