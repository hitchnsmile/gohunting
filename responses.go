package gohunting

// SearchResponse stores data from /domain-search
type SearchResponse struct {
	Data struct {
		Domain       string `json:"domain"`
		Webmail      bool   `json:"webmail"`
		Pattern      string `json:"pattern"`
		Organization string `json:"organization"`
		Emails       []struct {
			Value      string    `json:"value"`
			Type       string    `json:"type"`
			Confidence int       `json:"confidence"`
			Sources    []Sources `json:"sources"`
		} `json:"emails"`
	} `json:"data"`
	Meta struct {
		Results int `json:"results"`
		Limit   int `json:"limit"`
		Offset  int `json:"offset"`
		Params  struct {
			Domain     string `json:"domain"`
			Company    string `json:"company"`
			Type       string `json:"type"`
			Offset     int    `json:"offset"`
			Seniority  string `json:"seniority"`
			Department string `json:"department"`
		}
	}
}

// FindResponse stores data from /email-finder
type FindResponse struct {
	Data struct {
		Email   string    `json:"email"`
		Score   int       `json:"score"`
		Domain  string    `json:"domain"`
		Sources []Sources `json:"sources"`
	} `json:"data"`
	Meta struct {
		Params struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			FullName  string `json:"full_name"`
			Domain    string `json:"domain"`
			Company   string `json:"company"`
		} `json:"params"`
	} `json:"meta"`
}

// VerifyResponse stores data from /email-verifier
type VerifyResponse struct {
	Data struct {
		Result     string    `json:"result"`
		Score      int       `json:"score"`
		Email      string    `json:"email"`
		Regexp     bool      `json:"regexp"`
		Gibberish  bool      `json:"gibberish"`
		Disposable bool      `json:"disposable"`
		Webmail    bool      `json:"webmail"`
		MXRecords  bool      `json:"mx_records"`
		SMTPServer bool      `json:"smtp_server"`
		SMTPCheck  bool      `json:"smtp_check"`
		AcceptAll  bool      `json:"accept_all"`
		Sources    []Sources `json:"sources"`
	} `json:"data"`
	Meta struct {
		Params struct {
			Email string `json:"email"`
		}
	}
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
