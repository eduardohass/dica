package schemas

import (
	"time"

	"gorm.io/gorm"
)

type UserAnswer struct {
	gorm.Model
	IdQuestion int64
	IdAnswer1  int64
	IdAnswer2  int64
	IdAnswer3  int64
	IdAnswer4  int64
}

type UserAnswerResponse struct {
	IdQuestion uint      `json:"IdQuestion"`
	IdAnswer1  uint      `json:"idAnswer1"`
	IdAnswer2  uint      `json:"idAnswer2"`
	IdAnswer3  uint      `json:"idAnswer3"`
	IdAnswer4  uint      `json:"idAnswer4"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
}
