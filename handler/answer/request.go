package answer

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create answer
type CreateAnswerRequest struct {
	Answer string `json:"answer"`
}

func (r *CreateAnswerRequest) Validate() error {
	fmt.Println(r.Answer)
	if r.Answer == "" {
		return errParamIsRequired("answer", "string")
	}
	fmt.Println("DBG==passou pelas validações sem erro")
	return nil
}

// Update answer
type UpdateAnswerRequest struct {
	Answer string `json:"answer"`
}

func (r *UpdateAnswerRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.Answer != "" {
		fmt.Println("DBG==Answer = vazio")
		return nil
	}
	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
