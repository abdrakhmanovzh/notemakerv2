package service

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenTTL = 12 * time.Hour
)

type SignedDetails struct {
	Username string
	Uid      int `json:"uid"`
	jwt.StandardClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (string, error) {
	user.Password = hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", errors.New("username is not registered in the system")
	}

	passwordIsValid, _ := verifyPassword(password, user.Password)
	if !passwordIsValid {
		return "", errors.New("password is incorrect")
	}

	claims := &SignedDetails{
		Username: user.Username,
		Uid:      user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(tokenTTL).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Panic(err)
	}
	return token, err
}

func (s *AuthService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if ok && token.Valid {
		return claims.Uid, claims.Username, nil
	}
	return 0, "", errors.New("token claims are not of type *SignedDetails")
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword)
}

func verifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "email or password is incorrect"
		check = false
	}
	return check, msg
}
