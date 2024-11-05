package structs

type Response struct {
	Total    int64       `json:"total"`
	Items    []Item      `json:"items"`
}