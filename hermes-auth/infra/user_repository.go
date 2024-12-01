package infra

import (
    "context"
    "errors"
    "github.com/xu3cl40122/hermes/hermes-auth/models"
    "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

type MongoUserRepository struct {
    collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
    return &MongoUserRepository{
        collection: db.Collection("users"),
    }
}


func (r *MongoUserRepository) Create(ctx context.Context, user *models.User) error {
    _, err := r.collection.InsertOne(ctx, user)
    if err != nil {
        if mongo.IsDuplicateKeyError(err) {
            return errors.New("email already exists")
        }
        return err
    }
    return nil
}