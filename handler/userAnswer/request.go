package userAnswer

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create userAnswer
type CreateUserAnswerRequest struct {
	Id         int64 `json:"id"`
	IdQuestion int64 `json:"idQuestion"`
	IdAnswer1  int64 `json:"idAnswer1"`
	IdAnswer2  int64 `json:"idAnswer2"`
	IdAnswer3  int64 `json:"idAnswer3"`
	IdAnswer4  int64 `json:"idAnswer4"`
}

func (r *CreateUserAnswerRequest) Validate() error {
	if r.IdQuestion <= 0 && r.IdAnswer1 <= 0 && r.IdAnswer2 <= 0 && r.IdAnswer3 <= 0 && r.IdAnswer4 <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.IdQuestion <= 0 {
		return errParamIsRequired("idQuestion", "int64")
	}
	if r.IdAnswer1 <= 0 {
		return errParamIsRequired("idAnswer1", "int64")
	}
	if r.IdAnswer2 <= 0 {
		return errParamIsRequired("idAnswer2", "int64")
	}
	if r.IdAnswer3 <= 0 {
		return errParamIsRequired("idAnswer3", "int64")
	}
	if r.IdAnswer4 <= 0 {
		return errParamIsRequired("idAnswer4", "int64")
	}
	return nil
}

// Update userAnswer
type UpdateUserAnswerRequest struct {
	IdQuestion int64 `json:"idQuestion"`
	IdAnswer1  int64 `json:"idAnswer1"`
	IdAnswer2  int64 `json:"idAnswer2"`
	IdAnswer3  int64 `json:"idAnswer3"`
	IdAnswer4  int64 `json:"idAnswer4"`
}

func (r *UpdateUserAnswerRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.IdQuestion > 0 || r.IdAnswer1 > 0 || r.IdAnswer2 > 0 || r.IdAnswer3 > 0 || r.IdAnswer4 > 0 {
		return nil
	}
	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
