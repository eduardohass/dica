package letter

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create letter
type CreateLetterRequest struct {
	Type string `json:"type"`
}

func (r *CreateLetterRequest) Validate() error {
	fmt.Println(r.Type)
	if r.Type == "" {
		return errParamIsRequired("type", "string")
	}
	fmt.Println("DBG==passou pelas validações sem erro")
	return nil
}

// Update letter
type UpdateLetterRequest struct {
	Type string `json:"type"`
}

func (r *UpdateLetterRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.Type != "" {
		fmt.Println("DBG==TYPE = vazio")
		return nil
	}
	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
