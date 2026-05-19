# 開発セットアップ

このリポジトリの開発環境は Nix flakes と direnv を前提にしています。

`flake.nix` に Go、Bun、Python、uv、Buf、Protocol Buffers などの開発ツールを定義しているため、Nix と direnv を用意するとリポジトリごとに同じ開発環境を再現できます。

## 1. Nix をインストールする

Nix は Determinate Systems のインストーラで導入します。

```sh
curl --proto '=https' --tlsv1.2 -fsSL https://install.determinate.systems/nix | sh -s -- install --determinate
```

インストール後、シェルを開き直して `nix` コマンドが使えることを確認します。

```sh
nix --version
```

## 2. direnv と nix-direnv をインストールする

direnv と nix-direnv は Nix 経由でインストールします。

```sh
nix profile install nixpkgs#direnv nixpkgs#nix-direnv
```

## 3. direnv の hook を設定する

利用しているシェルに合わせて direnv の hook を追加します。

zsh の場合:

```sh
echo 'eval "$(direnv hook zsh)"' >> ~/.zshrc
exec zsh
```

bash の場合:

```sh
echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
exec bash
```

fish の場合:

```sh
echo 'direnv hook fish | source' >> ~/.config/fish/config.fish
exec fish
```

## 4. 開発環境を有効化する

リポジトリのルートに移動して direnv を許可します。

```sh
cd MuchUp-project
direnv allow
```

このリポジトリには `.envrc` があり、`use flake` によって `flake.nix` の dev shell が自動で読み込まれます。

## 5. 動作確認

以下のコマンドで、開発ツールが利用できることを確認します。

```sh
go version
bun --version
python --version
uv --version
buf --version
```

direnv を使わずに一時的に開発環境へ入る場合は、次のコマンドを使います。

```sh
nix develop
```

Python 依存は `llm-service/pyproject.toml` を source of truth として uv で管理します。

```sh
cd llm-service
uv sync
uv run python main.py
```

## トラブルシューティング

`.envrc` や `flake.nix` を変更した場合は、次のコマンドで direnv の環境を再読み込みします。

```sh
direnv reload
```

direnv が自動で有効にならない場合は、シェルの設定ファイルに hook が追加されているか確認してください。
