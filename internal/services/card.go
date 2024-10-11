package services

import (
	"card-validator/internal/entities"
	"card-validator/internal/errs"
	"strconv"
	"time"
)

type CardService struct{}

func NewCardService() *CardService {
	return &CardService{}
}

func (s *CardService) Validate(card entities.Card) error {
	if !s.luhnCheck(card.CardNumber) {
		return errs.ErrCardNumberInvalid
	}

	if card.ExpirationMonth < 1 || card.ExpirationMonth > 12 {
		return errs.ErrExpirationMonthInvalid
	}

	if card.ExpirationYear < time.Now().Year() {
		return errs.ErrExpiredYear
	}

	if s.isExpired(card.ExpirationMonth, card.ExpirationYear) {
		return errs.ErrExpiredMonth
	}

	return nil
}

func (s *CardService) luhnCheck(cardNumber string) bool {
	var sum int
	var alternate bool

	for i := len(cardNumber) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}

		if alternate {
			n *= 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}

		sum += n
		alternate = !alternate
	}

	return sum%10 == 0
}

func (s *CardService) isExpired(month, year int) bool {
	currentYear, currentMonth := time.Now().Year(), int(time.Now().Month())
	if year < currentYear || (year == currentYear && month < currentMonth) {
		return true
	}
	return false
}
