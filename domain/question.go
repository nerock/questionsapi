package domain

type Question struct {
	Text      string    `json:"text"`
	CreatedAt string    `json:"createdAt"`
	Choices   [3]Choice `json:"choices"`
}
