package lib

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
}

func SignJWT(userId uint, version int) (*Jwt, *CustomError) {
	accessTokenExp := time.Minute * 15
	refreshTokenExp := time.Hour * 24 * 7

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExp)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   "access_token",
		ID:        fmt.Sprintf("%d", userId),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExp)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    fmt.Sprintf("Version %d", version),
		Subject:   "refresh_token",
		ID:        fmt.Sprintf("%d", userId),
	})

	AccessToken, err := accessToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return nil, NewError("Failed to sign access token", 500, JwtService)
	}

	RefreshToken, err := refreshToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return nil, NewError("Failed to sign refresh token", 500, JwtService)
	}

	return &Jwt{
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
		CreatedAt:    time.Now(),
	}, nil
}

func DecodeJWT(token string) (jwt.MapClaims, *CustomError) {
	claims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return nil, NewError("Failed to decode token", 401, JwtService)
	}

	if !t.Valid {
		exp, err := claims.GetExpirationTime()
		if err != nil {
			return nil, NewError("Token is not valid", 401, JwtService)
		}
		expireIn := time.Now().Unix() - exp.Unix()
		if expireIn > 0 {
			return nil, NewError("Token has expired", 401, JwtService)
		} else {
			return nil, NewError("Token is not valid", 401, JwtService)
		}
	}

	return claims, &CustomError{}
}
