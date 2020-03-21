package model

type SearchResult struct {
	Href string `json:"href"`
	Title string `json:"title"`
	IsParse int64 `json:"is_parse"`
	Host string `json:"host"`
}

