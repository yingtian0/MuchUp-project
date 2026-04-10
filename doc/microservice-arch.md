               ┌──────────────┐
               │   Frontend   │
               │ (Web / SPA)  │
               └─────┬────────┘
                     │ REST / WebSocket
                     ▼
            ┌─────────────────────┐
            │   Envoy Gateway     │
            │  - Auth / Session   │
            │  - Rate Limit       │
            │  - HTTP → gRPC │
            │  - HTTP → WebSocket │
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

Observability / Monitoring:
┌───────────┐ ┌─────────────┐ ┌─────────────┐
│ Prometheus│ │ Grafana │ │ Jaeger / │
│ Metrics │ │ Dashboards │ │ OpenTelemetry│
└───────────┘ └─────────────┘ └─────────────┘

Notes / Data Flow (Draft):
- Primary message flow (sync):
  1. Client sends message via WebSocket.
  2. Envoy upgrades WS and forwards to API Service.
  3. API Service validates auth and publishes to Redis Streams.
  4. API Service writes durable record to Persistent DB.
  5. API Service pushes ack / broadcast to clients via WS.

- AI flow (async):
  1. API Service enqueues inference job to Redis Streams.
  2. AI Service consumes from Streams, runs model, returns result via gRPC or publishes to Streams.
  3. API Service persists AI output and pushes updates to clients.

- Ownership / scaling:
  - Envoy handles TLS termination, auth gate, rate limit, and WS upgrade.
  - API Service owns WS session state and must validate auth on every request.
  - Redis Streams is the durable queue/bus; Pub/Sub can be used only for non-critical fanout.
  - Persistent DB is the system of record (messages, sessions, AI results).
