package models

import (
	"time"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductName        string             `json:"pro_name" bson:"pro_name"`
	ProductDescription string             `json:"pro_description" bson:"pro_description"`
	TotalProduct       int64              `json:"total_pro" bson:"total_pro"`
	Price              float64            `json:"price"  bson:"price"`
	Colors             []string           `json:"colors" bson:"colors"`
	SellerId           primitive.ObjectID `json:"seller_id" bson:"seller_id"`
	CreatedAt          primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt          primitive.DateTime `json:"updated_at" bson:"updated_at" `
}

func (product *Products) SetProducts(moduleData dto.CreateProductDTO) Products {
	newProduct := Products{}
	sellerId, _ := helper.ConvertStringToPrimitive(moduleData.SellerId)

	newProduct.ProductName = moduleData.ProductName
	newProduct.ProductDescription = moduleData.ProductDescription
	newProduct.TotalProduct = moduleData.TotalProduct
	newProduct.Price = moduleData.Price
	newProduct.Colors = moduleData.Colors

	newProduct.SellerId = sellerId
	newProduct.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newProduct.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newProduct
}

type ProductImages struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductId  primitive.ObjectID `json:"prod_id" bson:"prod_id"`
	ProductImg string             `json:"prod_img" bson:"prod_img"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at" `
}

type Categories struct {
	Id                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CategoryType        string             `json:"cat_type" bson:"cat_type"`
	CategoryDescription string             `json:"cat_desc" bson:"cat_desc"`
	CreatedAt           primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt           primitive.DateTime `json:"updated_at" bson:"updated_at" `
}

func (cat *Categories) SetCategories(moduleData dto.CreateCategoriesDTO) Categories {
	newCategory := Categories{}

	newCategory.CategoryType = moduleData.CategoryType
	newCategory.CategoryDescription = moduleData.CategoryDescription
	newCategory.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newCategory.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newCategory
}

type ProductsSelling struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductId   primitive.ObjectID `json:"prod_id" bson:"prod_id"`
	CategoryId  primitive.ObjectID `json:"cat_id" bson:"cat_id"`
	IsAvailable bool               `json:"is_available" bson:"is_available"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

func (prodSell *ProductsSelling) SetProductsSelling(moduleData dto.CreateProductsSellingDTO) ProductsSelling {
	newProductSelling := ProductsSelling{}

	productID, _ := helper.ConvertStringToPrimitive(moduleData.ProductId)
	catID, _ := helper.ConvertStringToPrimitive(moduleData.CategoryId)

	newProductSelling.ProductId = productID
	newProductSelling.CategoryId = catID
	newProductSelling.IsAvailable = true
	newProductSelling.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	newProductSelling.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return newProductSelling
}
