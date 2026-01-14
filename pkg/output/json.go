package output

import (
	"encoding/json"
	"fmt"

	"github.com/safetest-dev/extractor-recon/internal/scanner"
)

func PrintJSON(results []scanner.Result) {
	data, _ := json.MarshalIndent(results, "", "  ")
	fmt.Println(string(data))
}
