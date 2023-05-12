package services

import (
	"errors"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/models"
	"github.com/aniket0951.com/product-service/repository"
)

type CategoryService interface {
	CreateCategory(category dto.CreateCategoriesDTO) error
	UpdateCategory(category dto.UpdateCategoriesDTO) error
	GetAllCategory() ([]models.Categories, error)
	CategoryById(catId string) (models.Categories, error)
	DeleteCategory(catId string) error
}

type categoryService struct {
	catRepo repository.CategoryRepository
}

func NewCategoryService(catRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		catRepo: catRepository,
	}
}

func (ser *categoryService) CreateCategory(category dto.CreateCategoriesDTO) error {

	err := ser.catRepo.CheckDuplicateCategory(category.CategoryType)
	if err != nil {
		return err
	}

	categoryToCreate := new(models.Categories).SetCategories(category)
	ser.catRepo.CreateCategory(categoryToCreate)
	return nil
}

func (ser *categoryService) UpdateCategory(category dto.UpdateCategoriesDTO) error {
	_, err := helper.ValidatePrimitiveId(category.Id)
	if err != nil {
		return err
	}
	return ser.catRepo.UpdateCategory(category)
}

func (ser *categoryService) GetAllCategory() ([]models.Categories, error) {
	return ser.catRepo.GetAllCategory()
}

func (ser *categoryService) CategoryById(catId string) (models.Categories, error) {
	catObjId, err := helper.ValidatePrimitiveId(catId)
	if err != nil {
		return models.Categories{}, err
	}

	res, resErr := ser.catRepo.CategoryById(catObjId)

	if resErr != nil {
		return models.Categories{}, resErr
	}
	if (res == models.Categories{}) {
		return models.Categories{}, errors.New("category not found for this id")
	}

	return res, nil
}
func (ser *categoryService) DeleteCategory(catId string) error {
	catObjId, err := helper.ValidatePrimitiveId(catId)
	if err != nil {
		return err
	}

	return ser.catRepo.DeleteCategory(catObjId)
}
