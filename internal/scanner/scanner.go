package scanner

import (
	"io"
	"net/http"
	"time"

	"github.com/username/extractor-recon/internal/parser"
	"github.com/username/extractor-recon/internal/probe"
)

type Scanner struct {
	Client     *http.Client
	StatusOnly bool
	Prober     *probe.Prober
}

func NewScanner(followRedirect bool, timeout int, statusOnly bool) *Scanner {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	if !followRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return &Scanner{
		Client:     client,
		StatusOnly: statusOnly,
		Prober:     probe.New(client),
	}
}

func (s *Scanner) Scan(target string) Result {
	resp, err := s.Client.Get(target)
	if err != nil {
		return Result{
			URL:   target,
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	result := Result{
		URL:        target,
		Status:     resp.StatusCode,
		StatusText: http.StatusText(resp.StatusCode),
	}

	if s.StatusOnly || resp.StatusCode >= 400 {
		return result
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Error = "failed to read response body"
		return result
	}

	links := parser.ExtractLinks(body, target)

	for _, link := range links {
		result.Links = append(result.Links, s.Prober.Check(link))
	}

	return result
}
