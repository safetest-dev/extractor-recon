package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/safetest-dev/extractor-recon/internal/config"
	"github.com/safetest-dev/extractor-recon/internal/scanner"
	"github.com/safetest-dev/extractor-recon/pkg/output"
)

func main() {
	cfg := parseFlags()

	sc := scanner.NewScanner(
		cfg.FollowRedirect,
		cfg.TimeoutSeconds,
		cfg.StatusOnly,
	)

	var results []scanner.Result
	for _, t := range cfg.Targets {
		results = append(results, sc.Scan(t))
	}

	if cfg.JSONOutput {
		output.PrintJSON(results)
	} else {
		output.PrintText(results)
	}
}

func parseFlags() config.Config {
	statusOnly := flag.Bool("status-only", false, "")
	jsonOut := flag.Bool("json", false, "")
	follow := flag.Bool("follow-redirect", false, "")
	timeout := flag.Int("timeout", 10, "")
	flag.Parse()

	return config.Config{
		StatusOnly:     *statusOnly,
		JSONOutput:     *jsonOut,
		FollowRedirect: *follow,
		TimeoutSeconds: *timeout,
		Targets:        loadTargets(flag.Args()),
	}
}

func loadTargets(args []string) []string {
	if len(args) != 1 {
		os.Exit(1)
	}

	arg := args[0]
	if strings.HasPrefix(arg, "http") {
		return []string{arg}
	}

	file, err := os.Open(arg)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var targets []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line != "" {
			targets = append(targets, line)
		}
	}
	return targets
}
