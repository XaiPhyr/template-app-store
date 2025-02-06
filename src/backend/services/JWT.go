package services

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	u "app_store_api/utils"
)

type JwtClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const (
	JwtExpiration        = 15 * time.Minute
	LimiterExpiration    = 1 * time.Minute
	JwtRefreshExpiration = 1 * time.Hour
)

var (
	JwtMethodHS256 = jwt.SigningMethodHS256
)

func GenerateJWT() (string, error) {
	app := u.InitConfig()
	jwtKey := app.Server.JwtKey

	claims := registerToken(JwtExpiration)

	token := jwt.NewWithClaims(JwtMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		log.Printf("Error: %s", err)
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(token string, w http.ResponseWriter) (err error) {
	app := u.InitConfig()
	jwtKey := app.Server.JwtKey

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	}

	t, err := jwt.ParseWithClaims(token, &JwtClaim{}, keyFunc)

	claims, _ := t.Claims.(*JwtClaim)

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Printf("Error: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	return
}

func RefreshJWT() {
	//
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Printf("Error: %s", err)
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	return nil
}

func registerToken(duration time.Duration) (claims *JwtClaim) {
	claims = &JwtClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	return
}
