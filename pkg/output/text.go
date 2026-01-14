package output

import (
	"fmt"

	"github.com/safetest-dev/extractor-recon/internal/scanner"
)

func PrintText(results []scanner.Result) {
	for _, r := range results {
		fmt.Printf("[+] %s\n", r.URL)
		fmt.Printf("[+] Status: %d %s\n", r.Status, r.StatusText)

		for _, l := range r.Links {
			if l.Error != "" {
				fmt.Printf("  [-] %s : ERROR (%s)\n", l.URL, l.Error)
			} else {
				fmt.Printf("  [-] %s : %d %s\n",
					l.URL, l.Status, l.StatusText)
			}
		}
	}
}
