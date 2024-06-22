package question

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create question
type CreateQuestionRequest struct {
	Question string `json:"question"`
}

func (r *CreateQuestionRequest) Validate() error {
	fmt.Println(r.Question)
	if r.Question == "" {
		return errParamIsRequired("question", "string")
	}
	fmt.Println("DBG==passou pelas validações sem erro")
	return nil
}

// Update question
type UpdateQuestionRequest struct {
	Question string `json:"question"`
}

func (r *UpdateQuestionRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.Question != "" {
		return nil
	}
	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
