package entity

type Album struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CoverUrl    string `json:"coverUrl"`
}
