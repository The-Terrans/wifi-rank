# wifi-rank

Runs a speed test, saves it under the current Wi-Fi name, and ranks every Wi-Fi you have tested by median download speed.

```text
#  WIFI          DOWN Mbps  UP Mbps  PING ms  TESTS
1  Office        137.4      49.1     4        3      <- you
2  Home          82.0       20.3     9        5
```

## Usage

```sh
wifi-rank         # test now, save, show ranking
wifi-rank -list   # show ranking only
```

History is stored as JSON lines in `~/.config/wifi-rank/history.jsonl`.

Requires Linux with NetworkManager (`nmcli`). Tests use Ookla speedtest.net servers via [speedtest-go](https://github.com/showwin/speedtest-go).

## Develop

```sh
nix develop -f shell.nix                 # Go from cache.nixos.org, no source build
go test ./...
CGO_ENABLED=0 go build -ldflags="-s -w"  # static binary, ~6 MB
```
