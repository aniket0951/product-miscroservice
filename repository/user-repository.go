package repository

import (
	"context"
	"os"

	"github.com/aniket0951.com/product-service/config"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userConnection        = config.GetCollection(config.DB, os.Getenv("USERS"))
	userAddressConnection = config.GetCollection(config.DB, os.Getenv("ADDRESS"))
)

type UserRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateUserAccount(models.Users) error
	GetUserByID(userId primitive.ObjectID) (interface{}, error)
	AddUserAddress(userAddress models.Address) error
	CheckDuplicateUser(email string) error
	CheckDuplicateAddress(userId primitive.ObjectID) error
}
type userRepository struct {
	userCollection        *mongo.Collection
	userAddressCollection *mongo.Collection
}

func NewUserRepository() UserRepository {
	return &userRepository{
		userCollection:        userConnection,
		userAddressCollection: userAddressConnection,
	}
}
