package main

import (
	"slices"
	"sort"
)

type ranking struct {
	SSID     string
	Tests    int
	Download float64
	Upload   float64
	PingMs   float64
}

// rank groups results by SSID and sorts by median download, fastest first.
// Medians keep one lucky or unlucky test from skewing a network's place.
func rank(results []result) []ranking {
	bySSID := map[string][]result{}
	for _, r := range results {
		bySSID[r.SSID] = append(bySSID[r.SSID], r)
	}

	var out []ranking
	for ssid, rs := range bySSID {
		out = append(out, ranking{
			SSID:     ssid,
			Tests:    len(rs),
			Download: median(rs, func(r result) float64 { return r.Download }),
			Upload:   median(rs, func(r result) float64 { return r.Upload }),
			PingMs:   median(rs, func(r result) float64 { return r.PingMs }),
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Download > out[j].Download })
	return out
}

func median(rs []result, field func(result) float64) float64 {
	vals := make([]float64, len(rs))
	for i, r := range rs {
		vals[i] = field(r)
	}
	slices.Sort(vals)
	mid := len(vals) / 2
	if len(vals)%2 == 0 {
		return (vals[mid-1] + vals[mid]) / 2
	}
	return vals[mid]
}
