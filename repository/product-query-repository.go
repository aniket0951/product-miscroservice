package repository

import (
	"context"
	"errors"
	"time"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *productRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}

func (db *productRepository) CreateProduct(product models.Products) {
	db.Product = append(db.Product, product)
}

func (db *productRepository) UpdateProduct(product dto.UpdateProductDTO) error {
	prodObjID, _ := helper.ConvertStringToPrimitive(product.Id)

	for i := range db.Product {
		if db.Product[i].Id == prodObjID {
			sellerID, _ := primitive.ObjectIDFromHex(product.Product.SellerId)
			db.Product[i].ProductName = product.Product.ProductName
			db.Product[i].ProductDescription = product.Product.ProductDescription
			db.Product[i].TotalProduct = product.Product.TotalProduct
			db.Product[i].Price = product.Product.Price
			db.Product[i].Colors = product.Product.Colors
			db.Product[i].SellerId = sellerID
			db.Product[i].UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

			return nil
		}
	}

	return errors.New("product not found to update")

}

func (db *productRepository) DeleteProduct(productId primitive.ObjectID) {

	for i := range db.Product {
		if db.Product[i].Id == productId {
			db.Product = append(db.Product[:i], db.Product[i+1:]...)
			break
		}
	}
}

func (db *productRepository) ProductsBySeller(sellerId primitive.ObjectID) []models.Products {
	res := []models.Products{}
	for i := range db.Product {
		if db.Product[i].SellerId == sellerId {
			res = append(res, db.Product[i])
		}
	}

	return res
}

func (db *productRepository) IncreaseTotalProduct(productId primitive.ObjectID, increase int) error {

	for i := range db.Product {
		if db.Product[i].Id == productId {
			db.Product[i].TotalProduct++
			return nil
		}
	}

	return errors.New("product not found to update")
}

func (db *productRepository) DecreaseTotalProduct(productId primitive.ObjectID, increase int) error {

	for i := range db.Product {
		if db.Product[i].Id == productId {
			db.Product[i].TotalProduct--
			return nil
		}
	}

	return errors.New("product not found to update")

}

func (db *productRepository) AddProductImg(productImg models.ProductImages) {
	db.ProductImages = append(db.ProductImages, productImg)
}

func (db *productRepository) ProductExitsOrNot(productId primitive.ObjectID) bool {
	for i := range db.Product {
		if db.Product[i].Id == productId {
			return true
		}
	}
	return false
}

// product for selling
func (db *productRepository) SellTheProduct(productSell models.ProductsSelling) {
	db.ProductSelling = append(db.ProductSelling, productSell)
}

func (db *productRepository) CheckProductAlreadyExits(productId primitive.ObjectID) error {

	for i := range db.ProductSelling {
		if db.ProductSelling[i].ProductId == productId {
			return errors.New("product is already in selling")
		}
	}

	return nil
}

func (db *productRepository) ProductsForSelling() (interface{}, error) {

	res := []models.ProductsForSellingResponseDTO{}

	for i := range db.Product {
		temp := models.ProductsForSellingResponseDTO{}
		temp.Products = db.Product[i]

		for j := range db.ProductSelling {
			if db.ProductSelling[j].ProductId == db.Product[i].Id {
				temp.Categories = db.GetProductCatagory(db.ProductSelling[j].CategoryId)
				break
			}
		}

		for k := range db.ProductImages {
			if db.ProductImages[k].Id == db.Product[i].Id {
				temp.ProductImages = db.GetProductImage(db.Product[i].Id)
				break
			}
		}

		res = append(res, temp)
	}
	return res, nil
}

func (db *productRepository) GetProductCatagory(catId primitive.ObjectID) models.Categories {
	for i := range db.Categories {
		if db.Categories[i].Id == catId {
			return db.Categories[i]
		}
	}

	return models.Categories{}
}

func (db *productRepository) GetProductImage(prodId primitive.ObjectID) models.ProductImages {
	for i := range db.ProductImages {
		if db.ProductImages[i].ProductId == prodId {
			return db.ProductImages[i]
		}
	}

	return models.ProductImages{}
}
