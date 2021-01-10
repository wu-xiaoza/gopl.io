package github

import "time"

//IssuesURL .
const IssuesURL = "https://api.github.com/search/issues"

//IssuesSearchResult .
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

//Issue .
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

//User .
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
