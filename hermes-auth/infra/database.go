package infra

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB 全局的 MongoDB 連接實例
var DB *mongo.Database

// ConnectDB 初始化 MongoDB 連接
func ConnectDB() {
	// 設置 MongoDB URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 連接到 MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// 設置數據庫名稱
	DB = client.Database("hermes")

	log.Println("Connected to MongoDB")
}
