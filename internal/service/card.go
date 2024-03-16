package service

import (
	"validate"
	"validate/internal/repository"
)

type CardService struct {
	repo repository.Card
}

func NewCardService(repo repository.Card) *CardService {
	return &CardService{repo: repo}
}

func (s *CardService) Create(number int, card validate.Card) (string, error) {
	return s.repo.Create(number, card)
}

func (s *CardService) GetAll(limit string, page string) ([]validate.Card, error) {
	return s.repo.GetAll(limit, page)

}

func (s *CardService) GetById(cardId string) (validate.Card, error) {
	return s.repo.GetById(cardId)

}

func (s *CardService) Delete(cardId string) error {
	return s.repo.Delete(cardId)
}

func (s *CardService) UpdateCard(cardId string, input validate.UpdateCardInput) error {
	return s.repo.UpdateCard(cardId, input)
}
