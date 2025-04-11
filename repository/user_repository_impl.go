package repository

import (
	"context"
	"password-manager/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}

}

func (r *userRepository) Create(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.db.Collection("user").InsertOne(ctx, u)

	return nil
}

func (r *userRepository) UpdateByUUID(uuid string, u *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := r.db.Collection("user").FindOne(ctx, uuid)
	r.db.Collection("user").DeleteOne(ctx, res)

	return u, nil

}

func (r *userRepository) DeleteByUUID(uuid string) error {
	return nil

}

func (r *userRepository) ListarUsers() ([]models.User, error) {

	return nil, nil

}
