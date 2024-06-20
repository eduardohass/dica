package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Letter struct {
	gorm.Model
	Type string
}

type LetterResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
