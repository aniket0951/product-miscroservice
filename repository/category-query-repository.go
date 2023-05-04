package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *categoryRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}

func (db *categoryRepository) CreateCategory(category models.Categories) error {
	ctx, cancel := db.Init()
	defer cancel()

	_, err := db.categoryCollection.InsertOne(ctx, category)
	return err
}

func (db *categoryRepository) UpdateCategory(category dto.UpdateCategoriesDTO) error {
	ctx, cancel := db.Init()
	defer cancel()

	catObjId, _ := helper.ConvertStringToPrimitive(category.Id)

	filter := bson.D{
		bson.E{Key: "_id", Value: catObjId},
	}

	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "cat_type", Value: category.CategoryType},
			bson.E{Key: "cat_desc", Value: category.CategoryDescription},
			bson.E{Key: "updated_at", Value: primitive.NewDateTimeFromTime(time.Now())},
		}},
	}

	res, err := db.categoryCollection.UpdateOne(ctx, filter, update)

	if res.MatchedCount == 0 {
		return errors.New("product not found to update")
	}

	if res.ModifiedCount == 0 {
		return errors.New("failed to update product details")
	}

	return err
}

func (db *categoryRepository) GetAllCategory() ([]models.Categories, error) {
	ctx, cancel := db.Init()
	defer cancel()

	cursor, curErr := db.categoryCollection.Find(ctx, bson.M{})
	if curErr != nil {
		return nil, curErr
	}

	var category []models.Categories

	if err := cursor.All(context.TODO(), &category); err != nil {
		return nil, err
	}

	return category, nil
}

func (db *categoryRepository) CategoryById(catId primitive.ObjectID) (models.Categories, error) {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "_id", Value: catId},
	}

	var category models.Categories
	err := db.categoryCollection.FindOne(ctx, filter).Decode(&category)
	return category, err
}

func (db *categoryRepository) DeleteCategory(catId primitive.ObjectID) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "_id", Value: catId},
	}
	res, err := db.categoryCollection.DeleteOne(ctx, filter)
	if res.DeletedCount == 0 {
		return errors.New("category not found to delete")
	}
	return err
}

func (db *categoryRepository) CheckDuplicateCategory(catType string) error {
	ctx, cancel := db.Init()
	defer cancel()

	filter := bson.D{
		bson.E{Key: "cat_type", Value: catType},
	}

	var category models.Categories

	err := db.categoryCollection.FindOne(ctx, filter).Decode(&category)
	fmt.Println("Error : ", err)
	if (category == models.Categories{}) {
		return nil
	}
	return errors.New("category already found")
}
