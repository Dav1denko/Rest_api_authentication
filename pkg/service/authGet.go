package service

import (
	restapiauthentication "Rest_api_authentication"
	"Rest_api_authentication/pkg/repository"
	"encoding/base64"
	"log"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	signingKey = "dfsdfsasdfsdfsdfsdfa"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GetTokens(GUID int) (restapiauthentication.Info, error) {
	Token, error := GenerateToken()
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

func GenerateToken() (string, error) {
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
