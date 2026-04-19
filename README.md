# MuchUp – AI が参加するリアルタイムグループチャット

MuchUp は、**5 人をランダムでマッチング**し、  
**AI が会話内容を理解して場の空気を良くする**ことを目的とした  
リアルタイムチャットアプリケーションです。


---

## コンセプト

- 5人ランダムマッチング
- AI が会話を解析し、盛り上げる

---

## 全体アーキテクチャ

```text
               ┌──────────────┐
               │   Frontend   │
               │ Web  │
               └─────┬────────┘
                     │ REST / WebSocket
                     ▼
            ┌─────────────────────┐
            │   Envoy Gateway     │
            │  - Auth / Session   │
            │  - Rate Limit       │
            │  - HTTP → WebSocket │
            │  - HTTP → gRPC │
            │  - TLS Termination  │
            └─────┬───────────────┘
                  │ gRPC / WebSocket
      ┌───────────┴───────────┐
      ▼                       ▼

┌──────┐ ┌───────┐
│ API Service │ │ AI Service │
│ - Business │ │ - AI / ML │
│ - WebSocket │ │ - gRPC / WS │
│ - Redis Pub/Sub / List │
└──────┘ └─────┘
│
▼
┌───────────────┐
│ Redis │
│ Cluster / │
│ Streams / Pub │
└───────────────┘
│
▼
┌───────────────┐
│ Persistent DB │
│ (Backup / Arch) │
└───────────────┘
```

- redis だけに room の状態を持たせる（単一 trust）
- envoy で gateway を定義してそこでリクエストをすべて認可・認証する。ビジネスロジックは外部を気にしない
