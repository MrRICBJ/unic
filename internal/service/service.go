package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
	"university/internal/entity"
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
	UserId   int64  `json:"id"`
	UserRole string `json:"role"`
}

func New(repo *repository.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTestAnswers() ([]string, error) {
	return s.repo.GetTestAnswers()
}

func (s *Service) GetTestQuestions() ([]string, error) {
	return s.repo.GetTestQuestions()
}

func (s *Service) GetMyStudent(id int64) ([]entity.User, error) {
	return s.repo.GetMyStudent(id)
}

func (s *Service) AddStudent(idStudent int, id int64) error {
	return s.repo.AddStudent(idStudent, id)
}

func (s *Service) CheckTest(id int64, answers []int) (int, error) {
	correctAnswers := []int{124, 134, 145, 1, 1, 2, 3, 1, 2, 1, 1, 14, 3}

	numCorrect := 0

	for i, userAnswer := range answers {
		if userAnswer == correctAnswers[i] {
			numCorrect++
		}
	}

	err := s.repo.CreateResultTest(numCorrect, id)
	if err != nil {
		return 0, err
	}

	return numCorrect, nil
}

func (s *Service) GetResultTests(id int64) ([]string, error) {
	return s.repo.GetResultTest(id)
}

func (s *Service) GetAllUsers() ([]entity.User, error) {
	return s.repo.GetUsers()
}

func (s *Service) GetTheory() (string, error) {
	return s.repo.GetTheory()
}

func (s *Service) CreateUser(user *entity.User) (*entity.User, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Service) GenerateToken(name, password string) (string, error) {
	user, err := s.repo.GetUser(name, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})

	token, err := tokenStr.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) ParseToken(accessToken string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.UserRole, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
