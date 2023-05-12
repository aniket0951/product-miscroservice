package repository

import (
	"context"
	"errors"
	"time"

	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *userRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}
func (db *userRepository) CreateUserAccount(userAccountData models.Users) {
	db.UserData = append(db.UserData, userAccountData)
}

func (db *userRepository) GetUserByID(userId primitive.ObjectID) (interface{}, error) {

	for i := range db.UserData {
		if db.UserData[i].ID == userId {
			return db.UserData[i], nil
		}
	}

	return nil, errors.New("user account not found")

}

func (db *userRepository) AddUserAddress(userAddress models.Address) {
	db.UserAddress = append(db.UserAddress, userAddress)
}

func (db *userRepository) CheckDuplicateAddress(userId primitive.ObjectID) error {

	for i := range db.UserData {
		if db.UserAddress[i].ID == userId {
			return errors.New("user account already found")
		}
	}
	return nil
}

func (db *userRepository) CheckDuplicateUser(email string) error {
	for i := range db.UserData {
		if db.UserData[i].Email == email {
			return errors.New("this email already have a account")
		}
	}
	return nil
}
