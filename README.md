# MuchUp – AI が参加するリアルタイムグループチャット

MuchUp は、**5 人をランダムでマッチング**し、  
**AI が会話内容を理解して場の空気を良くする**ことを目的とした  
リアルタイムチャットアプリケーションです。


---

## コンセプト

- 5人ランダムマッチング
- AI が会話を解析し、盛り上げる

---

## 開発セットアップ

開発環境は Nix flakes と direnv を前提にしています。

セットアップ手順は[docs/development/setup.md](docs/development/setup.md) を参照してください。

---

## 全体アーキテクチャ

```text
               ┌──────────────┐
               │   Frontend   │
               │     Web      │
               └─────┬────────┘
                     │ REST / WebSocket
                     ▼
        ┌─────────────────────┐        ┌─────────────────┐
        │   Envoy Gateway     │◀──────▶│  Auth Service    │
        │  - Auth / Session   │  gRPC  │  - Login/Auth    │
        │  - Rate Limit       │        │  - JWT / Session │
        │  - HTTP → WebSocket │        │  - Token Verify  │
        │  - HTTP → gRPC      │        │  - User Context  │
        │  - TLS Termination  │        └─────────────────┘
        └─────┬───────────────┘
              │ gRPC / WebSocket
              ▼

┌───────────────┐        ┌─────────────┐
│ API Service   │        │ AI Service  │
│ - Business    │  gRPC  │ - AI / ML   │
│ - WebSocket   │──────▶ │ - gRPC / WS │
│ - Redis client│        │             │
└───────┬───────┘        └─────────────┘
        │
        ▼
┌───────────────┐
│ Redis         │
│ - Hash        │
│ - Streams     │
└───────┬───────┘
        │
        ▼
┌───────────────┐
│ DB            │
└───────────────┘
```

- redis だけに room の状態を持たせる（単一 trust）
- envoy で gateway を定義してそこでリクエストをすべて認可・認証する。ビジネスロジックは外部を気にしない
