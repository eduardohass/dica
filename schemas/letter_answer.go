package schemas

import (
	"time"

	"gorm.io/gorm"
)

type LetterAnswer struct {
	gorm.Model
	IdLetter int64
	IdAnswer int64
}

type LetterAnswerResponse struct {
	IdLetter  uint      `json:"idLetter"`
	IdAnswer  uint      `json:"idAnswer"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
