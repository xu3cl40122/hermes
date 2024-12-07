package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
    "github.com/xu3cl40122/hermes/hermes-auth/models"
)

var jwtKey = []byte("smart_small_fat")

type Claims struct {
    Email string `json:"email"`
    ID string `json:"id"`
    jwt.RegisteredClaims
}

func GenerateJWT(user *models.User) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &Claims{
        Email: user.Email,
        ID: user.ID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }

    return claims, nil
}