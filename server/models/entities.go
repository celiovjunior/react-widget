package models

type Feedback struct {
	ID         int64  `json:"id"`
	Type       string `json:"type"`
	Comment    string `json:"comment"`
	Screenshot string `json:"screenshot"`
}
