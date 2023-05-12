package repository

import (
	"context"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateCategory(category models.Categories)
	UpdateCategory(category dto.UpdateCategoriesDTO) error
	GetAllCategory() ([]models.Categories, error)
	CategoryById(catId primitive.ObjectID) (models.Categories, error)
	DeleteCategory(catId primitive.ObjectID) error
	CheckDuplicateCategory(catType string) error
}
type categoryRepository struct {
	Categories []models.Categories
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		Categories: []models.Categories{},
	}
}
