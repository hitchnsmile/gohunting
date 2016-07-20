package gohunting

// SearchResponse stores data from /search
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

// FindResponse stores data from /generate
type FindResponse struct {
	Status  string    `json:"status"`
	Email   string    `json:"email"`
	Score   int       `json:"score"`
	Sources []Sources `json:"sources"`
}

// VerifyResponse stores data from /verify
type VerifyResponse struct {
	Status     string    `json:"status"`
	Email      string    `json:"email"`
	Result     string    `json:"result"`
	Score      int       `json:"score"`
	Regexp     bool      `json:"regexp"`
	Gibberish  bool      `json:"gibberish"`
	Disposable bool      `json:"disposable"`
	Webmail    bool      `json:"webmail"`
	MXRecords  bool      `json:"mx_records"`
	SMTPServer bool      `json:"smtp_server"`
	SMTPCheck  bool      `json:"smtp_check"`
	AcceptAll  bool      `json:"accept_all"`
	Sources    []Sources `json:"sources"`
}

// CountResponse stores data from /email-count
type CountResponse struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

// AccountResponse stores data from /account
type AccountResponse struct {
	Status    string `json:"status"`
	Email     string `json:"email"`
	PlanName  string `json:"plan_name"`
	PlanLevel int    `json:"plan_level"`
	ResetDate string `json:"reset_date"`
	Calls     struct {
		Used      int `json:"used"`
		Available int `json:"available"`
	} `json:"calls"`
}

// Sources stores data any endpoint that returns sources
type Sources struct {
	Domain      string `json:"domain"`
	URI         string `json:"uri"`
	ExtractedOn string `json:"extracted_on"`
}
