package schemas

import (
	"time"

	"gorm.io/gorm"
)

type UserAnswer struct {
	gorm.Model
	Id_question int64
	Id_answer_1 int64
	Id_answer_2 int64
	Id_answer_3 int64
	Id_answer_4 int64
}

type UserAnswerResponse struct {
	Id          uint      `json:"id"`
	Id_user     uint      `json:"id_user"`
	Id_question uint      `json:"id_question"`
	Id_answer   uint      `json:"id_answer"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}
