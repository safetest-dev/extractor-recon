package scanner

type LinkResult struct {
	URL        string `json:"url"`
	Status     int    `json:"status"`
	StatusText string `json:"status_text"`
	Error      string `json:"error,omitempty"`
}

type Result struct {
	URL        string       `json:"url"`
	Status     int          `json:"status"`
	StatusText string       `json:"status_text"`
	Links      []LinkResult `json:"links,omitempty"`
	Error      string       `json:"error,omitempty"`
}
