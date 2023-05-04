package repository

import (
	"context"
	"os"

	"github.com/aniket0951.com/product-service/config"
	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	productConnection        = config.GetCollection(config.DB, os.Getenv("PRODUCT"))
	productImgConnection     = config.GetCollection(config.DB, os.Getenv("PRODUCT_IMG"))
	productSellingConnection = config.GetCollection(config.DB, os.Getenv("PRODUCT_SELL"))
)

type ProductRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateProduct(product models.Products) error
	UpdateProduct(product dto.UpdateProductDTO) error
	DeleteProduct(productId primitive.ObjectID) error
	ProductsBySeller(sellerId primitive.ObjectID) ([]models.Products, error)
	IncreaseTotalProduct(productId primitive.ObjectID, increase int) error
	DecreaseTotalProduct(productId primitive.ObjectID, decrease int) error
	AddProductImg(productImg models.ProductImages) error
	ProductExitsOrNot(productId primitive.ObjectID) bool

	SellTheProduct(productSell models.ProductsSelling) error
	CheckProductAlreadyExits(productId primitive.ObjectID) error

	ProductsForSelling() ([]bson.M, error)
}
type productRepository struct {
	productCollection     *mongo.Collection
	productImgCollection  *mongo.Collection
	productSellCollection *mongo.Collection
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		productCollection:     productConnection,
		productImgCollection:  productImgConnection,
		productSellCollection: productSellingConnection,
	}
}
