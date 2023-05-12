package repository

import (
	"context"

	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateUserAccount(models.Users)
	GetUserByID(userId primitive.ObjectID) (interface{}, error)
	AddUserAddress(userAddress models.Address)
	CheckDuplicateUser(email string) error
	CheckDuplicateAddress(userId primitive.ObjectID) error
}
type userRepository struct {
	UserData    []models.Users
	UserAddress []models.Address
}

func NewUserRepository() UserRepository {
	return &userRepository{
		UserData:    []models.Users{},
		UserAddress: []models.Address{},
	}
}
