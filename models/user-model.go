package models

import (
	"strconv"
	"time"

	"github.com/aniket0951.com/product-service/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Contact   int64              `bson:"contact" json:"contact"`
	UserType  string             `bson:"user_type" json:"user_type"`
	Email     string             `json:"email" validate:"required"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at" `
}

func (user *Users) SetUsers(moduleData dto.CreateUserAccountDTO) Users {
	newUser := Users{}
	userContact, _ := strconv.Atoi(moduleData.Contact)
	newUser.ID = primitive.NewObjectID()
	newUser.Name = moduleData.Name
	newUser.Contact = int64(userContact)
	newUser.UserType = moduleData.UserType
	newUser.Email = moduleData.Email
	newUser.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newUser.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newUser
}

type Address struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AddressLine1 string             `json:"address_line_one" bson:"address_line_one"`
	AddressLine2 string             `json:"address_line_two" bson:"address_line_two"`
	City         string             `json:"city" bson:"city"`
	State        string             `json:"state" bson:"state"`
	PinCode      int                `json:"pin_code" bson:"pin_code"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt    primitive.DateTime `json:"updated_at" bson:"updated_at" `
}

func (address *Address) SetAddress(moduleData dto.CreateUserAddressDTO) Address {
	newAddress := Address{}

	userId, _ := primitive.ObjectIDFromHex(moduleData.UserID)
	newAddress.ID = primitive.NewObjectID()
	newAddress.AddressLine1 = moduleData.AddressLine1
	newAddress.AddressLine2 = moduleData.AddressLine2
	newAddress.City = moduleData.City
	newAddress.State = moduleData.State
	newAddress.PinCode = moduleData.PinCode
	newAddress.UserID = userId
	newAddress.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newAddress.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newAddress
}
