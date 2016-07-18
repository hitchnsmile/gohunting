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

type VerifyResponse struct {
	Status      string    `json:"status"`
	Email       string    `json:"email"`
	Result      string    `json:"result"`
	Score       int       `json:"score"`
	Regexp      bool      `json:"regexp"`
	Gibberish   bool      `json:"gibberish"`
	Disposable  bool      `json:"disposable"`
	Webmail     bool      `json:"webmail"`
	Mx_records  bool      `json:"mx_records"`
	Smtp_server bool      `json:"smtp_server"`
	Smtp_check  bool      `json:"smtp_check"`
	Accept_all  bool      `json:"accept_all"`
	Sources     []Sources `json:"sources"`
}

type CountResponse struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type AccountResponse struct {
	Status     string `json:"status"`
	Email      string `json:"email"`
	Plan_name  string `json:"plan_name"`
	Plan_level int    `json:"plan_level"`
	Reset_date string `json:"reset_date"`
	Calls      struct {
		Used      int `json:"used"`
		Available int `json:"available"`
	} `json:"calls"`
}

type Sources struct {
	Domain       string `json:"domain"`
	Uri          string `json:"uri"`
	Extracted_on string `json:"extracted_on"`
}
