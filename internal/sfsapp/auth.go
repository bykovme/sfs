package sfsapp

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims - jwt claims
type Claims struct {
	CustomID int64 `json:"user_id"`
	jwt.StandardClaims
}

func (app *App) CheckHeaderToken(r *http.Request) (err error) {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return errors.New("authorization token is missing")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return errors.New("authorization token has wrong format")
	}

	_, err = app.ValidateToken(splitToken[1])
	if err != nil {
		return err
	}
	return nil
}

// ValidateToken - validate if token is fine
func (app *App) ValidateToken(tokenString string) (userClaim Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(app.LoadedConfig.SignatureSalt), nil
	})
	if err != nil {
		return userClaim, err
	}
	if customClaim, ok := token.Claims.(*Claims); ok && token.Valid {
		userClaim.CustomID = customClaim.CustomID
		return userClaim, nil
	}
	return userClaim, errors.New("wrong token")
}

func (app *App) createToken(userID int64) (token string, err error) {
	expireToken := time.Now().Add(time.Hour * 24 * 365).Unix()
	claims := Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "sfs",
		},
	}
	// Create the token using your claims
	tokenEntity := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := tokenEntity.SignedString([]byte(app.LoadedConfig.SignatureSalt))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
