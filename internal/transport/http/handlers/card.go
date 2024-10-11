package handlers

import (
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
}

func NewCardHandler() *CardHandler {
	return &CardHandler{}
}

func (h *CardHandler) Register(router *gin.RouterGroup) {
	router.GET("", h.get)
}

func (h *CardHandler) get(ctx *gin.Context) {
	ctx.String(200, "12345467987")
}
