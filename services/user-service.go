package services

import (
	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/models"
	"github.com/aniket0951.com/product-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUserAccount(userAccountData dto.CreateUserAccountDTO) error
	GetUserByID(userId string) (interface{}, error)
	AddUserAddress(userAddress dto.CreateUserAddressDTO) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepository,
	}
}

func (ser *userService) CreateUserAccount(userAccountData dto.CreateUserAccountDTO) error {

	err := ser.userRepo.CheckDuplicateUser(userAccountData.Email)
	if err != nil {
		return err
	}

	newAccountCreate := new(models.Users).SetUsers(userAccountData)

	return ser.userRepo.CreateUserAccount(newAccountCreate)
}

func (ser *userService) GetUserByID(userId string) (interface{}, error) {

	userObj, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return models.Users{}, err
	}

	return ser.userRepo.GetUserByID(userObj)
}

func (ser *userService) AddUserAddress(userAddress dto.CreateUserAddressDTO) error {
	userObj, err := primitive.ObjectIDFromHex(userAddress.UserID)

	if err != nil {
		return err
	}
	dupErr := ser.userRepo.CheckDuplicateAddress(userObj)

	if dupErr != nil {
		return dupErr
	}

	addressToCreate := new(models.Address).SetAddress(userAddress)
	return ser.userRepo.AddUserAddress(addressToCreate)
}
