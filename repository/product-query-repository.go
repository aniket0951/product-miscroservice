package repository

import (
	"context"
	"errors"
	"time"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *productRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}

func (db *productRepository) CreateProduct(product models.Products) error {
	ctx, cancel := db.Init()
	defer cancel()
	_, err := db.productCollection.InsertOne(ctx, product)
	return err
}

func (db *productRepository) UpdateProduct(product dto.UpdateProductDTO) error {
	prodObjID, _ := helper.ConvertStringToPrimitive(product.Id)

	filter := bson.D{
		bson.E{Key: "_id", Value: prodObjID},
	}

	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "pro_name", Value: product.Product.ProductName},
			bson.E{Key: "pro_description", Value: product.Product.ProductDescription},
			bson.E{Key: "total_pro", Value: product.Product.TotalProduct},
			bson.E{Key: "price", Value: product.Product.Price},
			bson.E{Key: "colors", Value: product.Product.Colors},
			bson.E{Key: "seller_id", Value: product.Product.SellerId},
			bson.E{Key: "updated_at", Value: primitive.NewDateTimeFromTime(time.Now())},
		}},
	}

	ctx, cancel := db.Init()
	defer cancel()

	res, err := db.productCollection.UpdateOne(ctx, filter, update)

	if res.MatchedCount == 0 {
		return errors.New("product not found to update")
	}

	if res.ModifiedCount == 0 {
		return errors.New("failed to update product details")
	}

	return err
}

func (db *productRepository) DeleteProduct(productId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "_id", Value: productId},
	}

	_, err := db.productCollection.DeleteOne(ctx, filter)
	return err
}

func (db *productRepository) ProductsBySeller(sellerId primitive.ObjectID) ([]models.Products, error) {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "seller_id", Value: sellerId},
	}

	cursor, curErr := db.productCollection.Find(ctx, filter)

	if curErr != nil {
		return nil, curErr
	}

	var products []models.Products

	if err := cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (db *productRepository) IncreaseTotalProduct(productId primitive.ObjectID, increase int) error {
	filter := bson.D{
		bson.E{Key: "_id", Value: productId},
	}

	update := bson.D{
		bson.E{Key: "$inc", Value: bson.D{
			bson.E{Key: "total_pro", Value: increase},
		}},
	}

	ctx, cancel := db.Init()
	defer cancel()

	res, err := db.productCollection.UpdateOne(ctx, filter, update)

	if res.MatchedCount == 0 {
		return errors.New("product not found to update")
	}

	if res.ModifiedCount == 0 {
		return errors.New("failed to update product details")
	}

	return err
}

func (db *productRepository) DecreaseTotalProduct(productId primitive.ObjectID, increase int) error {
	filter := bson.D{
		bson.E{Key: "_id", Value: productId},
	}

	update := bson.D{
		bson.E{Key: "$inc", Value: bson.D{
			bson.E{Key: "total_pro", Value: -increase},
		}},
	}

	ctx, cancel := db.Init()
	defer cancel()

	res, err := db.productCollection.UpdateOne(ctx, filter, update)

	if res.MatchedCount == 0 {
		return errors.New("product not found to update")
	}

	if res.ModifiedCount == 0 {
		return errors.New("failed to update product details")
	}

	return err
}

func (db *productRepository) AddProductImg(productImg models.ProductImages) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.productImgCollection.InsertOne(ctx, &productImg)

	return err
}

func (db *productRepository) ProductExitsOrNot(productId primitive.ObjectID) bool {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "_id", Value: productId},
	}

	res := db.productCollection.FindOne(ctx, filter)

	return res.Err() == nil
}

// product for selling
func (db *productRepository) SellTheProduct(productSell models.ProductsSelling) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.productSellCollection.InsertOne(ctx, productSell)
	return err
}

func (db *productRepository) CheckProductAlreadyExits(productId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "prod_id", Value: productId},
	}

	var product models.ProductsSelling

	db.productSellCollection.FindOne(ctx, filter).Decode(&product)

	if (product == models.ProductsSelling{}) {
		return nil
	}

	return errors.New("product is already in selling")
}

func (db *productRepository) ProductsForSelling() ([]bson.M, error) {
	ctx, cancel := db.Init()
	defer cancel()

	filter := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "prod_id",
				"foreignField": "_id",
				"as":           "products",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "categories",
				"localField":   "cat_id",
				"foreignField": "_id",
				"as":           "category",
			},
		},

		{
			"$lookup": bson.M{
				"from":         "product_images",
				"localField":   "prod_id",
				"foreignField": "prod_id",
				"as":           "product_images",
			},
		},

		{"$unwind": "$products"},

		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "products.seller_id",
				"foreignField": "_id",
				"as":           "products.seller_info",
			},
		},

		{"$unwind": "$products.seller_info"},

		{
			"$lookup": bson.M{
				"from":         "user_address",
				"localField":   "user_id",
				"foreignField": "products.seller_id",
				"as":           "products.seller_info.seller_address",
			},
		},
	}

	cursor, curErr := db.productSellCollection.Aggregate(ctx, filter)

	if curErr != nil {
		return nil, curErr
	}

	var bdata []bson.M
	if err := cursor.All(context.TODO(), &bdata); err != nil {
		return nil, err
	}

	return bdata, nil
}
