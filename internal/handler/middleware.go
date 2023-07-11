package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	token := c.GetHeader(authorizationHeader)

	if token == "" {
		h.newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("handler - refreshToken - empty auth header "))
		return
	}
	headerParts := strings.Split(token, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		h.newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("handler - refreshToken - empty auth header"))
		return
	}

	if len(headerParts[1]) == 0 {
		h.newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("token is empty"))
		return
	}

	tokenValue, err := h.services.Users.ParseAccess(headerParts[1])
	if err != nil {
		h.newErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.Request = c.Request.WithContext(context.WithValue(context.Background(), "token", tokenValue))
	c.Next()
}
