# 0003. Use sqlc for data access layer

- Status: Proposed
- Date: 2026-06-15

## Context

<!--
なぜこの決定が必要になったのか。
前提、制約、課題を書く。
-->

- データアクセス層では、チャット送信、Room作成、履歴取得などのPostgreSQLクエリを実装する必要がある
- パフォーマンス要件としてDBクエリのp95レイテンシを意識する必要があり、発行されるSQLをレビューしやすいことが重要になる
- GoコードとSQLクエリの不整合をできるだけ早い段階で検出したい
- クエリ単位で計測、ログ、実行計画を確認しやすい構成にしたい

## Decision

<!--
何を決めたのか。
できるだけ短く、明確に書く。
-->

- データアクセス層のSQLコード生成ツールとしてsqlcを使用する。

## Rationale

<!--
なぜその決定にしたのか。
判断基準、重視したことを書く。
-->

- sqlcは手書きSQLからGoコードを生成するため、実際に実行されるSQLをレビュー対象にしやすい
- SQLを明示的に管理できるため、実行計画やインデックス設計をクエリ単位で確認しやすい
- 生成コードにより、SQLの入力パラメータや戻り値とGoコードの型の不整合をコンパイル時に検出しやすい
- クエリ名を単位として実装できるため、ログ、メトリクス、トレースなどの計装と対応づけやすい
- pgx向けのコード生成に対応しており、ADR 0002で採用したpgxと組み合わせやすい

## Consequences

<!--
この決定によって起きる影響を書く。
良い影響だけでなく、受け入れる制約や運用負荷も書く。
-->

- SQLファイルをデータアクセス層の主要なレビュー対象として扱う
- スキーマやSQLを変更した場合は、sqlcによるコード生成を実行して生成コードを更新する必要がある
- 複雑なドメイン操作では、SQL、生成コード、リポジトリ実装の責務分離を意識する必要がある
- ORMのようなリレーション操作や自動的なクエリ組み立ては提供されないため、必要なクエリは明示的にSQLとして定義する

## Alternatives Considered

<!--
検討したが採用しなかった選択肢を書く。
それぞれを不採用にした理由も書く。
-->

- ent
  - Goでテーブル定義を書き、生成されたClientやEntity APIを使ってDB操作を実装できる点は魅力がある
  - ただし、RoomとUserの関連取得などは生成APIを通して書くことになり、レビュー時に最初に見る対象がSQLではなくGoコードになる
  - 今回は発行されるSQLを先に確認し、クエリ単位で実行計画やインデックス設計を見たいので、entよりsqlcを優先した
- Bun
  - Goの構造体をテーブルに対応させつつ、Select、Join、Whereなどを使ってSQLに近い形でクエリを書ける
  - ただし、クエリはGoコードのメソッド呼び出しとして組み立てるため、SQLそのものを差分レビューする運用にはなりにくい
  - 今回はSQLファイルを主要なレビュー単位にしたいため、Bunよりsqlcを優先した
- GORM
  - 既存実装でも利用しており、構造体の保存、検索、関連データのPreloadなどを短いGoコードで書ける点は導入しやすい
  - ただし、Room作成時の関連保存や履歴取得時のPreloadなどは、実際にどのSQLが何回発行されるかをレビュー時に追いかける必要がある
  - 今回はSQLそのものをレビュー対象にして、クエリ単位で性能と計装を管理したいため、GORMよりsqlcを優先した

## References

<!--
参考資料、Issue、PR、公式ドキュメントなど。
-->

- [sqlc Documentation](https://docs.sqlc.dev/en/latest/)
- [sqlc - Using Go and pgx](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html)
- [ent Documentation](https://entgo.io/docs/getting-started/)
- [Bun Documentation](https://bun.uptrace.dev/guide/)
- [ADR 0002. Use pgx for postgres driver](./0002-use-pgx-for-postgres-driver.md)
