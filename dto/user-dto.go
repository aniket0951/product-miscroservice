package dto

import (
	"errors"
	"net/mail"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserAccountDTO struct {
	Name     string `json:"name" validate:"required"`
	Contact  string `json:"contact" validate:"required"`
	UserType string `json:"user_type" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func (userAccount *CreateUserAccountDTO) ValidateContact() error {
	num := userAccount.Contact
	msg := "please provide a valid contact number"
	if len(num) == 10 {
		_, err := strconv.Atoi(num)

		if err != nil {
			return errors.New(msg)
		}

		return nil

	} else {
		return errors.New(msg)
	}
}

func (userAccount *CreateUserAccountDTO) ValidateEmail() error {
	msg := "please provide a valid email address"

	_, err := mail.ParseAddress(userAccount.Email)

	if err != nil {
		return errors.New(msg)
	}

	return nil
}

type CreateUserAddressDTO struct {
	AddressLine1 string `json:"address_line_one" validate:"required"`
	AddressLine2 string `json:"address_line_two" validate:"required"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	UserID       string `json:"user_id" validate:"required"`
	PinCode      int    `json:"pin_code" validate:"required"`
}

func (address *CreateUserAddressDTO) ValidateUserID() error {
	_, err := primitive.ObjectIDFromHex(address.UserID)
	return err
}
