package main

import "testing"

func TestParseNmcli(t *testing.T) {
	out := "no:Neighbour:D8\\:85\\:5E\\:6C\\:25\\:86\n" +
		"yes:Cafe\\:Guest:D8\\:85\\:5E\\:5C\\:25\\:87\n"

	got, err := parseNmcli(out)
	if err != nil {
		t.Fatal(err)
	}
	want := wifi{SSID: "Cafe:Guest", BSSID: "D8:85:5E:5C:25:87"}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestParseNmcliNotConnected(t *testing.T) {
	if _, err := parseNmcli("no:Neighbour:D8\\:85\\:5E\\:6C\\:25\\:86\n"); err == nil {
		t.Error("expected error when no network is active")
	}
}

func TestRankUsesMedianDownload(t *testing.T) {
	results := []result{
		{SSID: "Home", Download: 50},
		{SSID: "Home", Download: 60},
		{SSID: "Home", Download: 900}, // one lucky spike
		{SSID: "Office", Download: 100},
		{SSID: "Office", Download: 120},
	}

	got := rank(results)

	if len(got) != 2 {
		t.Fatalf("got %d rankings, want 2", len(got))
	}
	if got[0].SSID != "Office" || got[0].Download != 110 || got[0].Tests != 2 {
		t.Errorf("first = %+v, want Office with median 110 over 2 tests", got[0])
	}
	if got[1].SSID != "Home" || got[1].Download != 60 {
		t.Errorf("second = %+v, want Home with median 60", got[1])
	}
}
