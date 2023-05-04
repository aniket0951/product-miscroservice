package controller

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController interface {
	CreateProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
	ProductsBySeller(*gin.Context)
	IncreaseDecreaseProduct(*gin.Context)

	AddProductImage(*gin.Context)

	AddProductForSell(*gin.Context)

	ProductsForSelling(*gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(prodSer services.ProductService) ProductController {
	return &productController{
		productService: prodSer,
	}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	productToCreate := dto.CreateProductDTO{}
	_ = ctx.BindJSON(&productToCreate)

	if (reflect.DeepEqual(productToCreate, dto.CreateProductDTO{})) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(&productToCreate); helper.CheckError(stErr, ctx) {
		return
	}

	_, valErr := helper.ValidatePrimitiveId(productToCreate.SellerId)

	if helper.CheckError(valErr, ctx) {
		return
	}

	err := c.productService.CreateProduct(productToCreate)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("product has been created", helper.EmptyObj{}, helper.PRODUCT_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	productToUpdate := dto.UpdateProductDTO{}
	_ = ctx.BindJSON(&productToUpdate)

	if (reflect.DeepEqual(productToUpdate, dto.UpdateProductDTO{})) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(&productToUpdate); helper.CheckError(stErr, ctx) {
		return
	}

	_, valErr := helper.ValidatePrimitiveId(productToUpdate.Id)
	if helper.CheckError(valErr, ctx) {
		return
	}

	err := c.productService.UpdateProduct(productToUpdate)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("product info has been updated", helper.EmptyObj{}, helper.PRODUCT_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	productId := helper.GetRequestQueryParam("product_id", ctx)

	if helper.CheckRequestParamEmpty(productId, ctx) {
		return
	}

	err := c.productService.DeleteProduct(productId)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.DELETE_SUCCESS, helper.EmptyObj{}, helper.PRODUCT_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) ProductsBySeller(ctx *gin.Context) {
	sellerID := helper.GetRequestQueryParam("seller_id", ctx)

	if helper.CheckRequestParamEmpty(sellerID, ctx) {
		return
	}

	res, err := c.productService.ProductsBySeller(sellerID)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, res, helper.PRODUCT_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) IncreaseDecreaseProduct(ctx *gin.Context) {
	tag := helper.GetRequestQueryParam("tag", ctx)
	number := helper.GetRequestQueryParam("number", ctx)
	productId := helper.GetRequestQueryParam("product_id", ctx)

	if helper.CheckRequestParamEmpty(tag, ctx) || helper.CheckRequestParamEmpty(number, ctx) || helper.CheckRequestParamEmpty(productId, ctx) {
		return
	}

	num, numErr := strconv.Atoi(number)

	if helper.CheckError(numErr, ctx) {
		return
	}

	switch {
	case tag == "increase":
		err := c.productService.IncreaseTotalProduct(productId, num)
		if helper.CheckError(err, ctx) {
			return
		}
		response := helper.BuildSuccessResponse("product has been increased", helper.EmptyObj{}, helper.PRODUCT_DATA)
		ctx.JSON(http.StatusOK, response)
		return
	case tag == "decrease":
		err := c.productService.DecreaseTotalProduct(productId, num)
		if helper.CheckError(err, ctx) {
			return
		}
		response := helper.BuildSuccessResponse("product has been decreased", helper.EmptyObj{}, helper.PRODUCT_DATA)
		ctx.JSON(http.StatusOK, response)
		return
	default:
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, "invalid tag received", helper.EmptyObj{}, helper.PRODUCT_DATA)
		ctx.JSON(http.StatusOK, response)

	}
}

func (c *productController) AddProductImage(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("image")
	productId := ctx.Request.FormValue("product_id")

	if helper.CheckError(err, ctx) {
		return
	}

	if helper.CheckRequestParamEmpty(productId, ctx) {
		return
	}

	resErr := c.productService.AddProductImg(productId, file)

	if helper.CheckError(resErr, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("product image added successfully", helper.EmptyObj{}, helper.CATEGORY_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) AddProductForSell(ctx *gin.Context) {
	productForSell := dto.CreateProductsSellingDTO{}
	_ = ctx.BindJSON(&productForSell)

	if (productForSell == dto.CreateProductsSellingDTO{}) {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	st := validator.New()

	if stErr := st.Struct(&productForSell); helper.CheckError(stErr, ctx) {
		return
	}

	err := c.productService.AddProductForSell(productForSell)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("product has been added for selling", helper.EmptyObj{}, helper.PRODUCT_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) ProductsForSelling(ctx *gin.Context) {
	res, err := c.productService.ProductsForSelling()

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, res, helper.PRODUCT_DATA)
	ctx.JSON(http.StatusOK, response)
}
