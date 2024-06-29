package internal

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
