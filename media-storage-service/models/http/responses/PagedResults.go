package responses

type PagedResult struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
