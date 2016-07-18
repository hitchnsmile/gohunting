package gohunting

type SearchResponse struct {
	Status  string `json:"status"`
	Domain  string `json:"domain"`
	Results int    `json:"results"`
	Webmail bool   `json:"webmail"`
	Pattern string `json:"pattern"`
	Offset  int    `json:"offset"`
	Emails  []struct {
		Value      string    `json:"value"`
		Type       string    `json:"type"`
		Confidence int       `json:"confidence"`
		Sources    []Sources `json:"sources"`
	} `json:"emails"`
}

type FindResponse struct {
	Status  string    `json:"status"`
	Email   string    `json:"email"`
	Score   int       `json:"score"`
	Sources []Sources `json:"sources"`
}

type Sources struct {
	Domain       string `json:"domain"`
	Uri          string `json:"uri"`
	Extracted_on string `json:"extracted_on"`
}
