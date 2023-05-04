package repository

import (
	"context"
	"errors"
	"time"

	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *userRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}
func (db *userRepository) CreateUserAccount(userAccountData models.Users) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.userCollection.InsertOne(ctx, userAccountData)

	return err
}

func (db *userRepository) GetUserByID(userId primitive.ObjectID) (interface{}, error) {
	ctx, cancel := db.Init()
	defer cancel()

	filter := []bson.M{

		{
			"$match": bson.M{
				"_id": userId,
			},
		},

		{
			"$lookup": bson.M{
				"from":         "user_address",
				"localField":   "_id",
				"foreignField": "user_id",
				"as":           "user_address",
			},
		},
	}

	cursor, curErr := db.userCollection.Aggregate(ctx, filter)

	if curErr != nil {
		return nil, curErr
	}

	var userData []bson.M
	if err := cursor.All(context.TODO(), &userData); err != nil {
		return nil, err
	}

	return userData, nil
}

func (db *userRepository) AddUserAddress(userAddress models.Address) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.userAddressCollection.InsertOne(ctx, userAddress)
	return err
}

func (db *userRepository) CheckDuplicateAddress(userId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "user_id", Value: userId},
	}
	var userAddress models.Address
	db.userAddressCollection.FindOne(ctx, filter).Decode(&userAddress)

	if (userAddress == models.Address{}) {
		return nil
	}
	return errors.New("user address already added")

}

func (db *userRepository) CheckDuplicateUser(email string) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "email", Value: email},
	}
	var userAddress models.Users
	db.userCollection.FindOne(ctx, filter).Decode(&userAddress)

	if (userAddress == models.Users{}) {
		return nil
	}
	return errors.New("this email already have a account")
}
