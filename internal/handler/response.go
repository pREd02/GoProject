package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message error) {
	h.logger.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		http.StatusText(statusCode),
	})
}
