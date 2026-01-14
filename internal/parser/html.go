package parser

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(body []byte, base string) []string {
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		return nil
	}

	seen := make(map[string]struct{})
	var results []string

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href := strings.TrimSpace(a.Val)
					if ignore(href) {
						continue
					}
					u, err := url.Parse(href)
					if err != nil {
						continue
					}
					full := baseURL.ResolveReference(u).String()
					if _, ok := seen[full]; !ok {
						seen[full] = struct{}{}
						results = append(results, full)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)
	return results
}

func ignore(h string) bool {
	return strings.HasPrefix(h, "#") ||
		strings.HasPrefix(h, "javascript:") ||
		strings.HasPrefix(h, "mailto:")
}
