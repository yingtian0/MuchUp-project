# 0002. Use pgx for postgres driver

- Status: Proposed
- Date: 2026-06-15

## Context

<!--
なぜこの決定が必要になったのか。
前提、制約、課題を書く。
-->

- PostgreSQLを利用する際に、コネクションから操作するまでのパッケージが必要になる
- パフォーマンス要件にもあるように、コネクションプールをチューニングできる必要がある
- チャット送信、Room作成、履歴取得ではPostgreSQLへのアクセスが発生するため、レイテンシと接続数を制御しやすいドライバを選ぶ必要がある
- PostgreSQLを前提にしたアプリケーションであるため、複数DBに共通する標準APIよりもPostgreSQL固有の機能を扱いやすいことを重視する

## Decision

<!--
何を決めたのか。
できるだけ短く、明確に書く。
-->

- PostgreSQLのDB Driverとしてpgxを使用する。

## Rationale

<!--
なぜその決定にしたのか。
判断基準、重視したことを書く。
-->

- pgxはPostgreSQL向けのドライバであり、PostgreSQL固有の型、エラー、機能を扱いやすいAPIが用意されている
- pgxpoolを利用することで、アプリケーション側でコネクションプールの最大接続数やタイムアウトを明示的に設定しやすい
- sqlcがpgx向けのコード生成に対応しており、SQLを中心にしたデータアクセス層と組み合わせやすい
- PostgreSQLを前提にするため、database/sqlの標準APIに寄せるよりも、pgxが提供する接続プールやPostgreSQL向けAPIを直接利用する方が今回の方針に合う

## Consequences

<!--
この決定によって起きる影響を書く。
良い影響だけでなく、受け入れる制約や運用負荷も書く。
-->

- コネクションプール、SQL実行、トランザクション管理をpgxのAPIに寄せて実装する
- sqlcで生成するコードもpgxを前提にできるため、データアクセス層の実装方針を揃えやすい
- PostgreSQLへの依存は強くなるため、将来別のRDBMSへ移行する場合はドライバと生成コードの見直しが必要になる
- pgx固有の型、エラー処理、接続管理をチームで理解して運用する必要がある

## Alternatives Considered

<!--
検討したが採用しなかった選択肢を書く。
それぞれを不採用にした理由も書く。
-->

- database/sql + lib/pq
  - Go標準のdatabase/sql APIに沿って実装でき、PostgreSQLドライバとしての利用実績もある
  - 標準APIに寄せられるため、接続、クエリ実行、トランザクションの扱いをGoの一般的なDB実装に合わせやすい
  - ただし、今回のアプリケーションはPostgreSQLを前提にしており、複数DBへの差し替えやすさよりもPostgreSQL向け機能の扱いやすさを重視する
  - pgxpoolによる接続プール設定、PostgreSQL固有の型やエラーの扱いやすさ、sqlcのpgx向けコード生成との相性を優先し、pgxを採用する

## References

<!--
参考資料、Issue、PR、公式ドキュメントなど。
-->

- [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)
- [sqlc - Using Go and pgx](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html)
- [lib/pq - Pure Go Postgres driver for database/sql](https://github.com/lib/pq)
