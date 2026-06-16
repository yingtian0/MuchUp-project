# Docker での起動手順

このドキュメントでは、リポジトリ直下の `compose.yaml` を使って Docker で開発用サービスを起動する手順を説明します。

## 前提

- Docker Desktop または Docker Engine がインストールされていること
- Docker Compose v2 の `docker compose` コマンドが使えること
- リポジトリのルートディレクトリでコマンドを実行すること

```sh
docker --version
docker compose version
```

## 起動されるサービス

現在の `compose.yaml` では、次のサービスを起動します。

| サービス | 内容 | 備考 |
| --- | --- | --- |
| `app` | Go の API サーバー | コンテナ内で HTTP `8080`、gRPC `9000` を listen |
| `db` | PostgreSQL 16 | ホスト側の `${DB_PORT}` に公開 |


## `.env` を作成する

`compose.yaml` は環境変数を参照するため、リポジトリ直下に `.env` を作成します。

```sh
touch .env
```

次の内容を `.env` に設定してください。

```env
DB_USER=muchup
DB_PASSWORD=muchup_password
DB_NAME=muchup
DB_PORT=5432
RUN_DB_MIGRATION=true
```

`RUN_DB_MIGRATION` は `compose.yaml` から `app` に渡されますが、現在の `app` 実装では起動時に常に GORM の `AutoMigrate` が実行されます。

## ビルドして起動する

```sh
docker compose up --build
```

バックグラウンドで起動する場合:

```sh
docker compose up --build -d
```

起動状況を確認します。

```sh
docker compose ps
```

ログを確認します。

```sh
docker compose logs -f app
docker compose logs -f db
```

## DB に接続する

`db` コンテナ内の `psql` で接続します。

```sh
docker compose exec db psql -U muchup -d muchup
```

ホストから接続する場合は、`.env` の `DB_PORT` を使います。

```sh
psql "postgres://muchup:muchup_password@localhost:5432/muchup?sslmode=disable"
```

`.env` の値を変更した場合は、上記の接続文字列も合わせて変更してください。

## 停止する

コンテナを停止します。

```sh
docker compose down
```

DB の永続データも削除して初期化する場合:

```sh
docker compose down -v
```

`-v` を付けると `postgres_data` volume が削除され、DB データは元に戻せません。

## トラブルシューティング

### `DB_USER` などが未定義というエラーが出る

リポジトリ直下に `.env` があるか確認してください。

```sh
ls -la .env
```

### `app` から DB に接続できない

`db` の healthcheck が通っているか確認してください。

```sh
docker compose ps
docker compose logs db
```

compose ネットワーク内では DB ホスト名は `db` です。`compose.yaml` の `app.environment.DB_HOST` は `db` のままにしてください。

### ホストから `localhost:8080` にアクセスできない

現在の `compose.yaml` は `app` のポートをホストに公開していません。ホストから API にアクセスする場合は、`app` サービスに次の設定を追加してから再起動してください。

```yaml
services:
  app:
    ports:
      - "8080:8080"
      - "9000:9000"
```

```sh
docker compose up --build
```

### Redis または LLM サービスへの接続で失敗する

`app` のデフォルト設定では、Redis は `localhost:6379`、LLM gRPC は `localhost:50052` を参照します。現在の `compose.yaml` には Redis と LLM サービスが含まれていないため、それらに依存する機能を Docker 上で動かす場合は、compose にサービスを追加し、`REDIS_ADDR` と `LLM_GRPC_ADDR` を `app.environment` に設定してください。
