package handlers

import (
	"card-validator/internal/entities"
	"card-validator/internal/errs"
	"card-validator/internal/transport/http/request"
	"errors"
)

func mapCardRequestToEntity(card request.CardRequest) entities.Card {
	return entities.Card{
		CardNumber:      card.CardNumber,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
	}
}

func mapErrorToCode(err error) int {
	switch {
	case errors.Is(err, errs.ErrCardNumberInvalid):
		return InvalidCardCode
	case errors.Is(err, errs.ErrExpirationMonthInvalid):
		return InvalidMonthCode
	case errors.Is(err, errs.ErrExpiredMonth):
		return ExpiredMonthCode
	case errors.Is(err, errs.ErrExpiredYear):
		return ExpiredYearCode
	default:
		return ValidationErrorCode
	}
}
