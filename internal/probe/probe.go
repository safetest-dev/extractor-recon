package probe

import (
	"net/http"

	"github.com/safetest-dev/extractor-recon/internal/scanner"
)

type Prober struct {
	Client *http.Client
}

func New(client *http.Client) *Prober {
	return &Prober{Client: client}
}

func (p *Prober) Check(url string) scanner.LinkResult {
	resp, err := p.Client.Get(url)
	if err != nil {
		return scanner.LinkResult{
			URL:   url,
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	return scanner.LinkResult{
		URL:        url,
		Status:     resp.StatusCode,
		StatusText: http.StatusText(resp.StatusCode),
	}
}
