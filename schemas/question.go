package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	IdQuestion int
	OptionD    string
	OptionI    string
	OptionC    string
	OptionA    string
}

type QuestionResponse struct {
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
	IdQuestion uint      `json:"idquestion"`
	OptionD    string    `json:"optiond"`
	OptionI    string    `json:"optioni"`
	OptionC    string    `json:"optionc"`
	OptionA    string    `json:"optiona"`
}
