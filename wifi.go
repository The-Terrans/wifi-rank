package main

import (
	"errors"
	"os/exec"
	"strings"
)

type wifi struct {
	SSID  string
	BSSID string
}

// currentWifi asks NetworkManager which access point is active.
func currentWifi() (wifi, error) {
	out, err := exec.Command("nmcli", "-t", "-f", "active,ssid,bssid", "dev", "wifi").Output()
	if err != nil {
		return wifi{}, errors.New("nmcli failed (is NetworkManager installed?): " + err.Error())
	}
	return parseNmcli(string(out))
}

func parseNmcli(out string) (wifi, error) {
	for _, line := range strings.Split(out, "\n") {
		fields := splitTerse(line)
		if len(fields) == 3 && fields[0] == "yes" {
			return wifi{SSID: fields[1], BSSID: fields[2]}, nil
		}
	}
	return wifi{}, errors.New("not connected to Wi-Fi")
}

// splitTerse splits an `nmcli -t` line on ':', honouring '\' escapes.
func splitTerse(line string) []string {
	var fields []string
	var cur strings.Builder
	escaped := false
	for _, r := range line {
		switch {
		case escaped:
			cur.WriteRune(r)
			escaped = false
		case r == '\\':
			escaped = true
		case r == ':':
			fields = append(fields, cur.String())
			cur.Reset()
		default:
			cur.WriteRune(r)
		}
	}
	return append(fields, cur.String())
}
