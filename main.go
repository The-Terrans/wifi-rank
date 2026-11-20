// wifi-rank runs a speed test, saves it under the current Wi-Fi name,
// and ranks every Wi-Fi you have tested by median download speed.
package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	list := flag.Bool("list", false, "show the ranking without running a new test")
	flag.Parse()

	if err := run(*list); err != nil {
		fmt.Fprintln(os.Stderr, "wifi-rank:", err)
		os.Exit(1)
	}
}

func run(listOnly bool) error {
	path, err := historyPath()
	if err != nil {
		return err
	}

	current := ""
	if !listOnly {
		w, err := currentWifi()
		if err != nil {
			return err
		}
		current = w.SSID

		fmt.Printf("Testing %q...\n", w.SSID)
		r, err := speedTest()
		if err != nil {
			return err
		}
		r.SSID, r.BSSID = w.SSID, w.BSSID
		fmt.Printf("Download %.1f Mbps | Upload %.1f Mbps | Ping %.0f ms | %s\n\n",
			r.Download, r.Upload, r.PingMs, r.Server)

		if err := appendHistory(path, r); err != nil {
			return err
		}
	}

	history, err := loadHistory(path)
	if err != nil {
		return err
	}
	if len(history) == 0 {
		fmt.Println("No tests yet. Run wifi-rank without -list.")
		return nil
	}
	printRanking(rank(history), current)
	return nil
}

func speedTest() (result, error) {
	servers, err := speedtest.New().FetchServers()
	if err != nil {
		return result{}, err
	}
	targets, err := servers.FindServer(nil)
	if err != nil {
		return result{}, err
	}
	s := targets[0]

	if err := s.PingTest(nil); err != nil {
		return result{}, err
	}
	if err := s.DownloadTest(); err != nil {
		return result{}, err
	}
	if err := s.UploadTest(); err != nil {
		return result{}, err
	}

	return result{
		Time:     time.Now(),
		Download: s.DLSpeed.Mbps(),
		Upload:   s.ULSpeed.Mbps(),
		PingMs:   float64(s.Latency.Microseconds()) / 1000,
		Server:   s.Sponsor + " (" + s.Name + ")",
	}, nil
}

func printRanking(rs []ranking, current string) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "#\tWIFI\tDOWN Mbps\tUP Mbps\tPING ms\tTESTS\t")
	for i, r := range rs {
		marker := ""
		if r.SSID == current {
			marker = "<- you"
		}
		fmt.Fprintf(tw, "%d\t%s\t%.1f\t%.1f\t%.0f\t%d\t%s\n",
			i+1, r.SSID, r.Download, r.Upload, r.PingMs, r.Tests, marker)
	}
	tw.Flush()
}
