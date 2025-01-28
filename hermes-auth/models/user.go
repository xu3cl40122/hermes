package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// User 定義用戶模型
type User struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	Nickname  string    `bson:"nickname" json:"nickname"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UserClaims struct {
	Email string `json:"email"`
	ID string `json:"id"`
	jwt.RegisteredClaims
}
