package letterAnswer

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create letterAnswer
type CreateLetterAnswerRequest struct {
	IdLetter int64 `json:"idLetter"`
	IdAnswer int64 `json:"idAnswer"`
}

func (r *CreateLetterAnswerRequest) Validate() error {
	if r.IdLetter <= 0 && r.IdLetter <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.IdLetter <= 0 {
		return errParamIsRequired("idLetter", "int64")
	}
	if r.IdLetter <= 0 {
		return errParamIsRequired("idLetter", "int64")
	}
	fmt.Println("DBG==passou pelas validações sem erro")
	return nil
}

// Update letterAnswer
type UpdateLetterAnswerRequest struct {
	IdLetter int64 `json:"idLetter"`
	IdAnswer int64 `json:"idAnswer"`
}

func (r *UpdateLetterAnswerRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.IdLetter > 0 || r.IdLetter > 0 {
		return nil
	}
	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
