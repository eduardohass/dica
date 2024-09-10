package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	IdAnswer   int
	IdUser     int
	IdQuestion int
	OptionD    bool
	OptionI    bool
	OptionC    bool
	OptionA    bool
}

type AnswerResponse struct {
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
	IdAnswer   uint      `json:"idanswer"`
	IdUser     int       `json:"iduser"`
	IdQuestion int       `json:"idquestion"`
	OptionD    bool      `json:"optiond"`
	OptionI    bool      `json:"optioni"`
	OptionC    bool      `json:"optionc"`
	OptionA    bool      `json:"optiona"`
}
