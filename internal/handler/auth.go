package handler

import (
	"GoProject/internal/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := h.services.Users.CreateUser(input)
	if err != nil {
		if errors.Is(err, models.ErrNotValidUser) {
			h.newErrorResponse(c, http.StatusBadRequest, err)
			return
		}
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, models.TokenValues{UserId: id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Auths struct {
	Auth   models.Auth   `json:"users"`
	Tokens models.Tokens `json:"tokens"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("delivery - signIn - BindJSON - %s", err.Error()))
		return
	}

	user, err := h.services.Users.GetUser(input.Username, input.Password)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("delivery - signIn - GetUser - %s", err.Error()))
		return
	}
	tokens, err := h.services.Users.GenerateToken(user.ID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("delivery - signIn - GenerateToken - %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, Auths{
		models.Auth{
			Username: user.Username},
		models.Tokens{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken},
	})
}

const (
	refresherHeader = "Refresher"
)

func (h *Handler) refresh(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	token := c.GetHeader(refresherHeader)
	if token == "" {
		h.newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("handler - refreshToken - empty auth header "))
		return
	}
	headerParts := strings.Split(token, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		h.newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("handler - refreshToken - empty auth header"))
		return
	}
	// Check refresh in black list

	tokens, err := h.services.Users.Refresh(headerParts[1])
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("handler - refresh - Refresh - %s", err.Error()))
		return
	}
	tokensvalue, err := h.services.Users.ParseAccess(tokens.AccessToken)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
	}
	user, err := h.services.Users.GetUserId(tokensvalue.UserId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	//	c.SetCookie("access", tokens.AccessToken, 360000000, "/", "", true, true)

	//	c.SetCookie("refresh", tokens.RefreshToken, 36000000000000, "/", "", true, true)
	c.JSON(http.StatusOK, Auths{
		models.Auth{
			Username: user.Username},
		models.Tokens{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken},
	})
	http.Redirect(c.Writer, c.Request, c.Request.Header.Get("Referer"), http.StatusSeeOther)
}
