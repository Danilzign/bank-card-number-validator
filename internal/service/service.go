package service

import (
	"validate"
	"validate/internal/repository"
)

type Card interface {
	Create(number int, card validate.Card) (string, error)
	GetAll(limit string, page string) ([]validate.Card, error)
	GetById(cardId string) (validate.Card, error)
	Delete(cardId string) error
	UpdateCard(cardId string, input validate.UpdateCardInput) error
}

type Service struct {
	Card
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Card: NewCardService(repos.Card),
	}
}
