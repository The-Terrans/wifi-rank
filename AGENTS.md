# AGENTS.md

## Develop

```sh
nix develop -f shell.nix                 # Go from cache.nixos.org, no source build
go test ./...
CGO_ENABLED=0 go build -ldflags="-s -w"  # static binary, ~6 MB
```
