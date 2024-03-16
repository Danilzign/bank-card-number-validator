package repository

import (
	"fmt"
	"strconv"
	"strings"
	"validate"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/theplant/luhn"
)

type CardPostgres struct {
	db *sqlx.DB
}

func NewCardPostgres(db *sqlx.DB) *CardPostgres {
	return &CardPostgres{db: db}
}

func (r *CardPostgres) Create(number int, card validate.Card) (string, error) {

	cardNumber := card.Number
	digitNumber, _ := strconv.Atoi(cardNumber)
	if digitNumber == 0 {
		return "", fmt.Errorf("Card number is not valid")
	}
	ok := luhn.Valid(digitNumber)

	var id string

	if ok {
		createCardQuery := "INSERT INTO cards (number) VALUES ($1) returning id"
		row := r.db.QueryRow(createCardQuery, card.Number)

		if err := row.Scan(&id); err != nil {
			return "", err
		}

	} else if !ok {
		return "", fmt.Errorf("Card number is not valid")
	}

	return id, nil
}

func (r *CardPostgres) GetById(cardId string) (validate.Card, error) {
	var card validate.Card

	query := "SELECT * FROM cards WHERE id=$1"
	if err := r.db.Get(&card, query, cardId); err != nil {
		return card, err
	}

	return card, nil
}

func (r *CardPostgres) GetAll(limit string, page string) ([]validate.Card, error) {

	var cards []validate.Card

	limitInt, _ := strconv.Atoi(limit)
	pageInt, _ := strconv.Atoi(page)
	offset := limitInt * (pageInt - 1)

	pagination := "SELECT * FROM cards ORDER BY id LIMIT $1 OFFSET $2"

	err := r.db.Select(&cards, pagination, limit, offset)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *CardPostgres) Delete(cardId string) error {

	query := "DELETE FROM cards WHERE id=$1"

	_, err := r.db.Exec(query, cardId)

	return err
}

func (r *CardPostgres) UpdateCard(cardId string, input validate.UpdateCardInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Number != nil {
		setValues = append(setValues, fmt.Sprintf("number=$%d", argId))
		args = append(args, *input.Number)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE cards SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, cardId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
