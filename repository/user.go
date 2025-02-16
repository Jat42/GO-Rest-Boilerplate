package repository

import (
	"context"
	"time"
	"users-rest/model"
	"users-rest/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}

func (r *UserRepository) Insert(user model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user.ID = util.NewUUID()
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []model.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) CreateIndexes() {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	r.collection.Indexes().CreateOne(context.Background(), indexModel)
}
