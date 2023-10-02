package models

type RequestInfo struct {
	IPAddress      string `json:"ip_address"`
	UserAgent      string `json:"user_agent"`
	Referrer       string `json:"referrer"`
	Timestamp      string `json:"timestamp"`
	AcceptLanguage string `json:"accept_language"`
	DoNotTrack     string `json:"do_not_track"`
}
