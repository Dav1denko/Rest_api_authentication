package service

import (
	restapiauthentication "Rest_api_authentication"
	"Rest_api_authentication/pkg/repository"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GetTokens(GUID int) (restapiauthentication.Info, error) {
	Token, error := GenerateNewToken()
	if error != nil {
		log.Printf("error generate token")
	}
	RefreshToken, error := GenerateNewRefreshToken()
	if error != nil {
		log.Printf("error generate new refresh token")
	}
	CreateGIUD := restapiauthentication.Info{
		GUID:         GUID,
		Token:        Token,
		RefreshToken: RefreshToken,
	}

	return CreateGIUD, nil
}

func GenerateNewToken() (string, error) {

	signingKey := os.Getenv("SIGNING_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(12 * time.Hour)},
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
	})
	return token.SignedString([]byte(signingKey))
}

func GenerateNewRefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}
	RefreshTokenWithBase64 := base64.StdEncoding.EncodeToString(b)
	return RefreshTokenWithBase64, nil
}

func (s *AuthService) SaveRefreshToken(GUID int, RefreshToken string) {
	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(RefreshToken), bcrypt.DefaultCost)
	if err != nil {
		fmt.Sprintln("error hashed token")
	}
	s.repo.SaveTokens(GUID, hashedRefreshToken)
}

func (s *AuthService) RefreshTokens(GUID int, CookieRefreshToken string) (bool, error) {
	if s.repo.GetRefreshTokens(GUID, CookieRefreshToken) {
		return true, nil
	}

	return false, nil
}
