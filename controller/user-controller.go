package controller

import (
	"net/http"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	CreateUserAccount(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	AddUserAddress(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userSer services.UserService) UserController {
	return &userController{
		userService: userSer,
	}
}

func (c *userController) CreateUserAccount(ctx *gin.Context) {
	userData := dto.CreateUserAccountDTO{}
	_ = ctx.BindJSON(&userData)

	if (userData == dto.CreateUserAccountDTO{}) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(userData); stErr != nil {
		helper.BuildUnProcessableEntity(ctx, stErr)
		return
	}

	if helper.CheckError(userData.ValidateContact(), ctx) || helper.CheckError(userData.ValidateEmail(), ctx) {
		return
	}

	err := c.userService.CreateUserAccount(userData)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.DATA_INSERTED, helper.EmptyObj{}, helper.USER_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	defer ctx.Request.Body.Close()
	userId := helper.GetRequestQueryParam("user_id", ctx)

	if helper.CheckRequestParamEmpty(userId, ctx) {
		return
	}

	res, err := c.userService.GetUserByID(userId)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, res, helper.USER_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) AddUserAddress(ctx *gin.Context) {
	userAddress := dto.CreateUserAddressDTO{}
	_ = ctx.BindJSON(&userAddress)

	if (userAddress == dto.CreateUserAddressDTO{}) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(&userAddress); helper.CheckError(stErr, ctx) {
		return
	}

	if helper.CheckError(userAddress.ValidateUserID(), ctx) {
		return
	}

	err := c.userService.AddUserAddress(userAddress)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.DATA_INSERTED, helper.EmptyObj{}, helper.USER_DATA)
	ctx.JSON(http.StatusOK, response)
}
