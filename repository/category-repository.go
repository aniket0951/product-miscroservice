package repository

import (
	"context"
	"os"

	"github.com/aniket0951.com/product-service/config"
	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	categoryConnection = config.GetCollection(config.DB, os.Getenv("CATEGORY"))
)

type CategoryRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateCategory(category models.Categories) error
	UpdateCategory(category dto.UpdateCategoriesDTO) error
	GetAllCategory() ([]models.Categories, error)
	CategoryById(catId primitive.ObjectID) (models.Categories, error)
	DeleteCategory(catId primitive.ObjectID) error
	CheckDuplicateCategory(catType string) error
}
type categoryRepository struct {
	categoryCollection *mongo.Collection
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		categoryCollection: categoryConnection,
	}
}
