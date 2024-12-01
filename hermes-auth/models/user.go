package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User 定義用戶模型
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // MongoDB 的 ObjectID
	Email    string             `bson:"email"`         // 用戶郵箱
	Password string             `bson:"password"`      // 密碼（需要加密存儲）
	Nickname string             `bson:"nickname"`      // 昵称
}
