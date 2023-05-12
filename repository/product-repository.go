package repository

import (
	"context"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateProduct(product models.Products)
	UpdateProduct(product dto.UpdateProductDTO) error
	DeleteProduct(productId primitive.ObjectID)
	ProductsBySeller(sellerId primitive.ObjectID) []models.Products
	IncreaseTotalProduct(productId primitive.ObjectID, increase int) error
	DecreaseTotalProduct(productId primitive.ObjectID, decrease int) error
	AddProductImg(productImg models.ProductImages)
	ProductExitsOrNot(productId primitive.ObjectID) bool

	SellTheProduct(productSell models.ProductsSelling)
	CheckProductAlreadyExits(productId primitive.ObjectID) error

	ProductsForSelling() (interface{}, error)
}
type productRepository struct {
	Product        []models.Products
	ProductImages  []models.ProductImages
	ProductSelling []models.ProductsSelling
	Categories     []models.Categories
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		Product:        []models.Products{},
		ProductImages:  []models.ProductImages{},
		ProductSelling: []models.ProductsSelling{},
		Categories:     []models.Categories{},
	}
}
