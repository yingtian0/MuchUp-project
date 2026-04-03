# MuchUp – AI が参加するリアルタイムグループチャット

MuchUp は、**5 人をランダムでマッチング**し、  
**AI が会話内容を理解して場の空気を良くする**ことを目的とした  
リアルタイムチャットアプリケーションです。

Envoy を API Gateway としたマイクロサービス構成により、  
高いスケーラビリティ・リアルタイム性・可観測性を備えています。

---

## コンセプト

- 5 人ランダムマッチング
- WebSocket によるリアルタイムチャット
- AI が会話を解析し、空気を和らげる / 盛り上げる
- API Gateway 集約型の認証・制御
- Observability による運用前提設計

---

## 全体アーキテクチャ

```text
┌────────────────────────────┐
│        Browser (SPA)        │
│  - HTTPS (REST/JSON)        │
│  - WebSocket                │
└─────────────┬──────────────┘
              │
              v
┌──────────────────────────────────────────┐
│           Envoy API Gateway               │
│  - JWT verification                       │
│  - Rate limit / Routing                   │
│  - WebSocket upgrade                      │
└─────────────┬────────────────────────────┘
              │
              │ HTTPS / REST
              v
┌──────────────────────────────────────────┐
│          Main API / Backend API           │
│                  (Go)                     │
│  - Business logic                         │
│  - Room/tenant authorization              │
│  - Session orchestration                  │
│  - REST endpoint for browser              │
└─────────────┬────────────────────────────┘
              │
              │ gRPC
              v
┌──────────────────────────────────────────┐
│               AI Service                  │
│                (Python)                   │
│  - Prompt assembly                        │
│  - Tool execution                         │
│  - LLM request normalization              │
└─────────────┬────────────────────────────┘
              │
              │ gRPC or HTTPS
              v
┌──────────────────────────────────────────┐
│               LLM Service                 │
│         (internal or external)            │
└──────────────────────────────────────────┘

                   + Redis
          - ephemeral room/session state
          - streams / pubsub / cache
```

- redis だけに room の状態を持たせる（単一 trust）
- envoy で gateway を定義してそこでリクエストをすべて認可・認証する。ビジネスロジックは外部を気にしない
