{
  description = "MuchUp development toolchain";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };

        python = pkgs.python312.withPackages (
          ps: with ps; [
            django
            grpcio
            protobuf
          ]
        );

        go = pkgs.go_1_26;
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            gotools
            golangci-lint

            bun

            python

            buf
            protobuf
            protoc-gen-go
            protoc-gen-go-grpc
            grpcurl
          ];

          shellHook = ''
            export BUF_CACHE_DIR="''${BUF_CACHE_DIR:-$PWD/.cache/buf}"
            export GOPATH="''${GOPATH:-$PWD/.cache/go}"
            export GOBIN="$GOPATH/bin"
            export PATH="$GOBIN:$PATH"

            mkdir -p "$BUF_CACHE_DIR" "$GOBIN"

            echo "MuchUp dev shell"
            echo "  go:     $(go version)"
            echo "  bun:    $(bun --version)"
            echo "  python: $(python --version)"
            echo "  buf:    $(buf --version)"
          '';
        };
      }
    );
}
