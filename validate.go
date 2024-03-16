package validate

import "errors"

type Card struct {
	Id     string `json:"id" db:"id"`
	Number string `json:"number" db:"number"`
}

type UpdateCardInput struct {
	Id     *string `json:"id"`
	Number *string `json:"number"`
}

func (i UpdateCardInput) Validate() error {
	if i.Id == nil && i.Number == nil {
		return errors.New("No values")
	}
	return nil
}
