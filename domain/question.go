package domain

import "time"

type Question struct {
	Text      string
	CreatedAt time.Time
	Choices   []Choice
}
