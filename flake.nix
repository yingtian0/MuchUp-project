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

        go = pkgs.go_1_26;

        hkVersion = "1.46.0";
        hk = pkgs.rustPlatform.buildRustPackage rec {
          doCheck = false;
          pname = "hk";
          version = hkVersion;

          src = pkgs.fetchFromGitHub {
            owner = "jdx";
            repo = "hk";
            rev = "v${version}";
            hash = "sha256-nuWanHAMgaL24CsOXPs8qFVrSDWDLkzJTKLHG/VQvKc=";
          };

          cargoHash = "sha256-TVDM94HFB39ys6zoaD+AgJvoMDMe4xigFAKpRITsZ1k=";

          meta = {
            description = "Git hook manager";
            homepage = "https://github.com/jdx/hk";
            license = pkgs.lib.licenses.mit;
            mainProgram = "hk";
          };

          nativeBuildInputs = [
            pkgs.pkg-config
            pkgs.pkl
          ];
          buildInputs = [
            pkgs.openssl
          ];
        };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            gotools
            golangci-lint

            bun

            python312
            uv

            buf
            protobuf
            protoc-gen-go
            protoc-gen-go-grpc
            grpcurl

            hk
            pkl
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
            echo "  uv:     $(uv --version)"
            echo "  buf:    $(buf --version)"
          '';
        };
      }
    );
}
