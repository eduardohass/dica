package answer

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create answer
type CreateAnswerRequest struct {
	IdAnswer   int  `json:"idanswer"`
	IdUser     int  `json:"iduser"`
	IdQuestion int  `json:"idquestion"`
	OptionD    bool `json:"optiond"`
	OptionI    bool `json:"optioni"`
	OptionC    bool `json:"optionc"`
	OptionA    bool `json:"optiona"`
}

func (r *CreateAnswerRequest) Validate() error {
	if r.IdAnswer == 0 {
		return errParamIsRequired("idanswer", "int")
	}
	if r.IdUser == 0 {
		return errParamIsRequired("iduser", "int")
	}
	if r.IdQuestion == 0 {
		return errParamIsRequired("idquestion", "int")
	}
	fmt.Println("DBG==passou pelas validações sem erro")
	return nil
}

// Update answer
type UpdateAnswerRequest struct {
	IdAnswer   int  `json:"idanswer"`
	IdUser     int  `json:"iduser"`
	IdQuestion int  `json:"idquestion"`
	OptionD    bool `json:"optiond"`
	OptionI    bool `json:"optioni"`
	OptionC    bool `json:"optionc"`
	OptionA    bool `json:"optiona"`
}

func (r *UpdateAnswerRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.IdAnswer == 0 && r.IdUser == 0 && r.IdQuestion == 0 {
		fmt.Println("DBG==tudo vazio")
		return nil
	}
	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
