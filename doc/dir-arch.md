y-chat-app/
├── buf.yaml
├── buf.gen.yaml
├── go.work
├── docker-compose.yaml
├── Makefile
│
├── proto/                       # 【真実の源】ここだけ人間が編集
│   └── chat/
│       └── v1/
│           ├── chat.proto
│           └── auth.proto
│
├── gen/                         # 【自動生成エリア】手動編集禁止
│   ├── go/                      # Backend 用 (types, stubs)
│   │   └── chat/v1/
│   ├── ts/                      # Frontend 用 (client & types)
│   └── openapi/                 # Gateway/Docs 用 (openapi.yaml)
│
├── backend/                     # 【Go monolith service】
│   ├── go.mod
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   └── internal/
│       ├── domain/
│       │   ├── model/           # User, Room, Message など
│       │   └── repository/      # インターフェース群
│       ├── usecase/
│       │   ├── matching.go      # 5人マッチング
│       │   ├── chat.go          # publish/subscribe, fanout, AI判定
│       │   └── auth.go          # token validate（ext_authz用もここ）
│       ├── adapter/
│       │   ├── handler/         # HTTP/WS ハンドラ (DTO<->Domain)
│       │   ├── presenter/       # レスポンス整形（必要なら）
│       │   └── gateway/
│       │       ├── redis_pubsub.go  # Pub/Sub 実装（subscribe/publish）
│       │       └── auth_store.go    # セッション管理（Redis or inmem）
│       └── ai/
│           ├── analyzer.go      # “AIはライブラリ”：判定API
│           └── providers/       # OpenAI等の実装差し替え層（任意）
│
├── frontend/
│   ├── package.json             # "@api": "file:../gen/ts"
│   ├── vite.config.ts
│   └── src/
│       ├── main.tsx
│       ├── hooks/
│       └── components/
│
├── gateway/                     # 【Envoy Proxy】（残すなら）
│   ├── Dockerfile
│   └── envoy.yaml               # gen/openapi/openapi.yaml を参照
│
└── k8s/
