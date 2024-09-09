package question

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create question
type CreateQuestionRequest struct {
	IdQuestion int    `json:"idquestion"`
	OptionD    string `json:"optiond"`
	OptionI    string `json:"optioni"`
	OptionC    string `json:"optionc"`
	OptionA    string `json:"optiona"`
}

func (r *CreateQuestionRequest) Validate() error {
	if r.IdQuestion == 0 {
		return errParamIsRequired("idquestion", "int")
	}
	if r.OptionD == "" && r.OptionI == "" && r.OptionC == "" && r.OptionA == "" {
		return errParamIsRequired("option", "string")
	}
	fmt.Println("DBG==passou pelas validações sem erro")
	return nil
}

// Update question
type UpdateQuestionRequest struct {
	IdQuestion int    `json:"idquestion"`
	OptionD    string `json:"optiond"`
	OptionI    string `json:"optioni"`
	OptionC    string `json:"optionc"`
	OptionA    string `json:"optiona"`
}

func (r *UpdateQuestionRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.IdQuestion != 0 && r.OptionD != "" && r.OptionI != "" && r.OptionC != "" && r.OptionA != "" {
		return nil
	}
	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
