package structs

type Response struct {
	Href     string      `json:"href"`
	Limit    int64       `json:"limit"`
	Next     string      `json:"next"`
	Offset   int64       `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int64       `json:"total"`
	Items    []Item      `json:"items"`
}