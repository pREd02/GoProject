package service

import (
	"GoProject/internal/models"
	repository "GoProject/internal/reposistory"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const (
	tokenAT = 30 * 24 * time.Hour
	tokenRT = 30 * 24 * time.Hour
)

type Users interface {
	CreateUser(user models.User) (int, error)
	ParseAccess(tokenPars string) (models.TokenValues, error)
	Refresh(tokenPars string) (models.Tokens, error)
	GenerateToken(id int) (models.Tokens, error)
	GetUser(username, password string) (models.User, error)
	GetUserId(id int) (models.User, error)
}

type UserService struct {
	repo repository.Users
}

func NewUser(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(user models.User) (int, error) {

	user.Password = generatePasswordHash(user.Password)
	return u.repo.CreateUser(user)
}

func (u *UserService) GetUser(username, password string) (models.User, error) {
	return u.repo.GetUser(username, generatePasswordHash(password))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"uuid"`
}

func (u *UserService) ParseAccess(tokenPars string) (models.TokenValues, error) {
	tokenParse, err := jwt.ParseWithClaims(tokenPars, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("signing_key")), nil
	})
	if err != nil {
		return models.TokenValues{}, models.ErrorTokenTime
	}
	claim, ok := tokenParse.Claims.(*tokenClaims)
	if !ok {
		return models.TokenValues{}, errors.New("token claims are not of type *tokenClaims")
	}
	var tokenValue models.TokenValues
	tokenValue.UserId = claim.UserID
	return tokenValue, nil
}
func (u *UserService) Refresh(tokenPars string) (models.Tokens, error) {

	token, err := jwt.ParseWithClaims(tokenPars, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("signing_key")), nil
	})
	if err != nil {
		return models.Tokens{}, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return models.Tokens{}, errors.New("token claims are not of type *tokenClaims")
	}
	tokens, err := u.GenerateToken(claims.UserID)
	if err != nil {
		return models.Tokens{}, nil
	}
	return tokens, nil
}

func newRefreshToken(userId int) (string, error) {
	claims := tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenRT).Unix(),
		},
		userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("signing_key")))
}

func newAccessToken(userId int) (string, error) {
	claims := tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenAT).Unix(),
		},
		userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("signing_key")))
}
func (u *UserService) GetUserId(id int) (models.User, error) {
	return u.repo.GetUserId(id)
}

func (u *UserService) GenerateToken(id int) (models.Tokens, error) {
	var tokens models.Tokens
	var err error
	tokens.AccessToken, err = newAccessToken(id)
	if err != nil {
		return models.Tokens{}, err
	}
	tokens.RefreshToken, err = newRefreshToken(id)
	if err != nil {
		return models.Tokens{}, err
	}
	return tokens, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("salt"))))
}
