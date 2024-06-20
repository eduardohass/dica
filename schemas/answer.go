package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	Id     int64
	Answer string
}

type AnswerResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Answer    string    `json:"answer"`
}
