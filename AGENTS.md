# AGENTS.md

## Develop

```sh
nix develop -f shell.nix                 # Go from cache.nixos.org, no source build
go test ./...
CGO_ENABLED=0 go build -ldflags="-s -w"  # static binary, ~6 MB
```

## Commits

Use [Conventional Commits](https://www.conventionalcommits.org/): `type: summary`, lowercase, imperative, no trailing period.

| Type | Use for |
|---|---|
| `feat` | new behavior |
| `fix` | bug fix |
| `test` | adding or changing tests |
| `docs` | README, AGENTS.md, comments |
| `chore` | tooling, deps, module setup |

Example: `feat: rank networks by median download speed`
