my-chat-app/
├── buf.yaml # Buf モジュール設定 (依存関係定義)
├── buf.gen.yaml # コード生成レシピ (Go, TS, Python, OpenAPI を一括生成)
├── go.work # Go Workspace (backend と gen/go を繋ぐ重要ファイル)
├── docker-compose.yaml # ローカル開発用 (Redis, Envoy, DB などを起動)
├── Makefile # "make gen", "make run" 等のタスクランナー
│
├── proto/ # 【真実の源】ここだけを人間が編集する
│ └── chat/
│ └── v1/
│ ├── chat.proto # チャット・マッチング定義
│ └── auth.proto # 認証定義
│
├── gen/ # 【自動生成エリア】手動編集禁止 (.gitignore 推奨だが開発中は入れても良い)
│ ├── go/ # Backend 用 (gRPC server stubs & types)
│ │ └── chat/v1/
│ ├── ts/ # Frontend 用 (Connect-ES Client & Types)
│ ├── python/ # AI Service 用 (gRPC code)
│ └── openapi/ # Envoy/Docs 用 (openapi.yaml)
│
├── backend/ # 【Go Service】クリーンアーキテクチャ採用
│ ├── go.mod # gen/go をローカル参照する
│ ├── cmd/
│ │ └── server/
│ │ └── main.go # エントリーポイント (DI と起動処理)
│ └── internal/ # 外部から隠蔽されたアプリケーションコード
│ ├── domain/ # [Entity] 依存なし・純粋な Go 構造体と IF 定義
│ │ ├── model/ # User, Message などのドメインモデル
│ │ └── repository/ # ChatRepo, AuthRepo などのインターフェース
│ ├── usecase/ # [Logic] ビジネスロジック
│ │ ├── matching_usecase.go # 5 人マッチングのロジック
│ │ └── chat_usecase.go # メッセージ処理・AI 介入判定依頼
│ └── adapter/ # [Interface Adapter] 外部との変換層
│ ├── handler/ # gRPC ハンドラー (Proto 型 <-> Domain 型 変換)
│ └── gateway/ # DB/外部 API への接続実装
│ ├── redis_repo.go # Redis への読み書き実装
│ └── ai_client.go # AI Service への gRPC Client 実装
│
├── frontend/ # 【React SPA】
│ ├── package.json # "dependencies": { "@api": "file:../gen/ts" }
│ ├── vite.config.ts
│ └── src/
│ ├── main.tsx
│ ├── hooks/ # 生成されたクライアントを使うカスタムフック
│ └── components/ # UI コンポーネント
│
├── ai-service/ # 【Python Service】
│ ├── pyproject.toml # Poetry or requirements.txt
│ ├── main.py # gRPC Server 起動
│ └── service/
│ └── chat_analyzer.py # LLM を使った空気読みロジック
│
├── gateway/ # 【Envoy Proxy】
│ ├── Dockerfile
│ └── envoy.yaml # gen/openapi/openapi.yaml を元にルーティング設定
│
└── k8s/ # 【Infrastructure】
├── redis/ # Redis Cluster Manifests
├── backend.yaml
├── frontend.yaml
├── ai-service.yaml
└── envoy.yaml
