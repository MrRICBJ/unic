package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
	"university/internal/entity"
	"university/internal/handlers"
	"university/internal/repository"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type Service struct {
	repo *repository.Repo
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"id"`
}

func New(repo *repository.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTheory() (string, error) {
	return s.repo.GetTheory()
}

func (s *Service) CreateUser(user *entity.User) (*entity.User, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Service) GenerateToken(name, password string) (*handlers.ResponseSignIn, error) {
	var result handlers.ResponseSignIn
	var err error

	result.User, err = s.repo.GetUser(name, generatePasswordHash(password))
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		result.User.Id,
	})

	result.Token, err = token.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Service) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
