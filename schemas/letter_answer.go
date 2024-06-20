package schemas

import (
	"time"

	"gorm.io/gorm"
)

type LetterAnswer struct {
	gorm.Model
	Id_letter int64
	Id_answer int64
}

type LetterAnswerResponse struct {
	Id_letter uint      `json:"id_letter"`
	Id_answer uint      `json:"id_answer"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
