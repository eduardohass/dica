package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Question string
}

type QuestionResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Question  string    `json:"question"`
}
