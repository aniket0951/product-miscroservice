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

func (db *categoryRepository) Init() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}

func (db *categoryRepository) CreateCategory(category models.Categories) {
	db.Categories = append(db.Categories, category)
}

func (db *categoryRepository) UpdateCategory(category dto.UpdateCategoriesDTO) error {
	catObjId, _ := helper.ConvertStringToPrimitive(category.Id)
	for i := range db.Categories {
		if db.Categories[i].Id == catObjId {
			db.Categories[i].CategoryType = category.CategoryType
			db.Categories[i].CategoryDescription = category.CategoryDescription
			db.Categories[i].UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
			return nil
		}
	}

	return errors.New("product not found to update")
}

func (db *categoryRepository) GetAllCategory() ([]models.Categories, error) {
	return db.Categories, nil
}

func (db *categoryRepository) CategoryById(catId primitive.ObjectID) (models.Categories, error) {

	for i := range db.Categories {
		if db.Categories[i].Id == catId {
			return db.Categories[i], nil
		}
	}
	return models.Categories{}, errors.New("category not found")
}

func (db *categoryRepository) DeleteCategory(catId primitive.ObjectID) error {
	for i := range db.Categories {
		if db.Categories[i].Id == catId {
			db.Categories = append(db.Categories[:i], db.Categories[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found to delete")
}

func (db *categoryRepository) CheckDuplicateCategory(catType string) error {

	for i := range db.Categories {
		if db.Categories[i].CategoryType == catType {
			return errors.New("category already found")
		}
	}

	return nil
}
