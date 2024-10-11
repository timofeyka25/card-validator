package handlers

import (
	"card-validator/internal/services"
	"card-validator/internal/transport/http/request"
	"card-validator/internal/transport/http/response"
	"card-validator/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ParsingErrorCode    = 1
	ValidationErrorCode = 2
	InvalidCardCode     = 3
	InvalidMonthCode    = 4
	ExpiredMonthCode    = 5
	ExpiredYearCode     = 6
)

type CardHandler struct {
	cardService *services.CardService
	validator   *validator.Validator
}

func NewCardHandler(cardService *services.CardService, validator *validator.Validator) *CardHandler {
	return &CardHandler{
		cardService: cardService,
		validator:   validator,
	}
}

func (h *CardHandler) Register(router *gin.RouterGroup) {
	card := router.Group("card")

	card.POST("validate", h.validate)
}

// @Summary Validate a credit card
// @Description Validates a credit card based on number, expiration month, and expiration year
// @Tags Card
// @Accept  json
// @Produce  json
// @Param card body request.CardRequest true "Card details to validate"
// @Success 200 {object} response.CardResponse "Validation result"
// @Failure 400 {object} response.CardResponse "Validation error"
// @Failure 422 {object} response.CardResponse "Parsing error"
// @Router /card/validate [post]
func (h *CardHandler) validate(ctx *gin.Context) {
	var card request.CardRequest

	if err := ctx.ShouldBindJSON(&card); err != nil {
		h.respondWithError(ctx, ParsingErrorCode, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.validator.Validate(card); err != nil {
		h.respondWithError(ctx, ValidationErrorCode, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.cardService.Validate(mapCardRequestToEntity(card)); err != nil {
		code := mapErrorToCode(err)
		h.respondWithError(ctx, code, err.Error(), http.StatusOK)
		return
	}

	h.respondWithSuccess(ctx)
}

func (h *CardHandler) respondWithError(ctx *gin.Context, code int, message string, status int) {
	ctx.JSON(status, response.CardResponse{
		Valid: false,
		Error: &response.CardError{
			Code:    code,
			Message: message,
		},
	})
}

func (h *CardHandler) respondWithSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.CardResponse{
		Valid: true,
		Error: nil,
	})
}
