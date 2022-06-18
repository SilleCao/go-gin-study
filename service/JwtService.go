package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IJWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JWTCustomsClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type JWTService struct {
	secretKey string
	issuer    string
}

func NewJWTService() IJWTService {
	return &JWTService{
		secretKey: getSecretKey(),
		issuer:    "sille.cn",
	}
}

func getSecretKey() string {
	secret := "0000"
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (jwtSrv *JWTService) GenerateToken(username string, admin bool) string {
	claims := &JWTCustomsClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtSrv *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
