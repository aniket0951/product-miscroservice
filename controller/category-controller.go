package controller

import (
	"net/http"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoryController interface {
	CreateCategory(*gin.Context)
	UpdateCategory(*gin.Context)
	GetAllCategory(*gin.Context)
	CategoryById(*gin.Context)
	DeleteCategory(*gin.Context)
}

type categoryController struct {
	catService services.CategoryService
}

func NewCategoryController(catSer services.CategoryService) CategoryController {
	return &categoryController{
		catService: catSer,
	}
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	categoryToCreate := dto.CreateCategoriesDTO{}
	_ = ctx.BindJSON(&categoryToCreate)

	if (categoryToCreate == dto.CreateCategoriesDTO{}) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(&categoryToCreate); helper.CheckError(stErr, ctx) {
		return
	}

	err := c.catService.CreateCategory(categoryToCreate)
	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("category has been created", helper.EmptyObj{}, helper.CATEGORY_DATA)
	ctx.JSON(http.StatusOK, response)

}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	categoryToUpdate := dto.UpdateCategoriesDTO{}
	_ = ctx.BindJSON(&categoryToUpdate)

	if (categoryToUpdate == dto.UpdateCategoriesDTO{}) {
		helper.RequestBodyEmptyResponse(ctx)
	}

	st := validator.New()

	if stErr := st.Struct(&categoryToUpdate); helper.CheckError(stErr, ctx) {
		return
	}

	err := c.catService.UpdateCategory(categoryToUpdate)
	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("category has been updated", helper.EmptyObj{}, helper.CATEGORY_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *categoryController) GetAllCategory(ctx *gin.Context) {
	res, err := c.catService.GetAllCategory()
	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, res, helper.CATEGORY_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *categoryController) CategoryById(ctx *gin.Context) {
	catId := helper.GetRequestQueryParam("cat_id", ctx)

	if helper.CheckRequestParamEmpty(catId, ctx) {
		return
	}

	res, err := c.catService.CategoryById(catId)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, res, helper.CATEGORY_DATA)
	ctx.JSON(http.StatusOK, response)

}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	catId := helper.GetRequestQueryParam("cat_id", ctx)

	if helper.CheckRequestParamEmpty(catId, ctx) {
		return
	}

	err := c.catService.DeleteCategory(catId)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.DELETE_SUCCESS, helper.EmptyObj{}, helper.CATEGORY_DATA)
	ctx.JSON(http.StatusOK, response)

}
