package repository

import (
	"validate"

	"github.com/jmoiron/sqlx"
)

type Card interface {
	Create(number int, card validate.Card) (string, error)
	GetAll(limit string, page string) ([]validate.Card, error)
	GetById(cardId string) (validate.Card, error)
	Delete(cardId string) error
	UpdateCard(cardId string, input validate.UpdateCardInput) error
}

type Repository struct {
	Card
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Card: NewCardPostgres(db),
	}
}
