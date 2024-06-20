package schemas

import (
	"time"

	"gorm.io/gorm"
)

type QuestionAnswer struct {
	gorm.Model
	Id_question int64
	Id_answer_1 int64
	Id_answer_2 int64
	Id_answer_3 int64
	Id_answer_4 int64
}

type QuestionAnswerResponse struct {
	Id_question uint      `json:"id_question"`
	Id_answer_1 uint      `json:"id_answer_1"`
	Id_answer_2 uint      `json:"id_answer_2"`
	Id_answer_3 uint      `json:"id_answer_3"`
	Id_answer_4 uint      `json:"id_answer_4"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}
