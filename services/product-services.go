package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/aniket0951.com/product-service/dto"
	"github.com/aniket0951.com/product-service/helper"
	"github.com/aniket0951.com/product-service/models"
	"github.com/aniket0951.com/product-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	CreateProduct(product dto.CreateProductDTO) error
	UpdateProduct(product dto.UpdateProductDTO) error
	DeleteProduct(productId string) error
	ProductsBySeller(sellerId string) ([]models.Products, error)
	IncreaseTotalProduct(productId string, increase int) error
	DecreaseTotalProduct(productId string, decrease int) error
	AddProductImg(productId string, file multipart.File) error

	AddProductForSell(productSell dto.CreateProductsSellingDTO) error
	CheckProductAlreadyExits(productId string) error

	ProductsForSelling() (interface{}, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepository,
	}
}

func (ser *productService) CreateProduct(product dto.CreateProductDTO) error {
	_, err := helper.ValidatePrimitiveId(product.SellerId)
	if err != nil {
		return err
	}
	newProductToCreate := new(models.Products).SetProducts(product)
	ser.productRepo.CreateProduct(newProductToCreate)
	return nil
}

func (ser *productService) UpdateProduct(product dto.UpdateProductDTO) error {
	return ser.productRepo.UpdateProduct(product)
}

func (ser *productService) DeleteProduct(productId string) error {
	productObjId, err := helper.ValidatePrimitiveId(productId)
	if err != nil {
		return err
	}
	ser.productRepo.DeleteProduct(productObjId)
	return nil
}
func (ser *productService) ProductsBySeller(sellerId string) ([]models.Products, error) {
	sellerObjId, err := helper.ValidatePrimitiveId(sellerId)
	if err != nil {
		return nil, err
	}
	res := ser.productRepo.ProductsBySeller(sellerObjId)
	return res, nil
}
func (ser *productService) IncreaseTotalProduct(productId string, increase int) error {
	productObjId, err := helper.ValidatePrimitiveId(productId)
	if err != nil {
		return err
	}
	return ser.productRepo.IncreaseTotalProduct(productObjId, increase)
}
func (ser *productService) DecreaseTotalProduct(productId string, decrease int) error {
	productObjId, err := helper.ValidatePrimitiveId(productId)
	if err != nil {
		return err
	}
	return ser.productRepo.DecreaseTotalProduct(productObjId, decrease)
}

func (ser *productService) AddProductImg(productId string, file multipart.File) error {
	productObjId, proErr := helper.ValidatePrimitiveId(productId)
	if proErr != nil {
		return proErr
	}

	if !ser.productRepo.ProductExitsOrNot(productObjId) {
		return errors.New("product doesn't exits to add a image")
	}

	tempFile, err := ioutil.TempFile("static", "upload-*.png")
	if err != nil {
		return err
	}

	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(tempFile)

	fileBytes, fileReaderr := ioutil.ReadAll(file)

	if fileReaderr != nil {
		return fileReaderr
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			panic(err)
		}
	}(tempFile)

	productImg := models.ProductImages{
		ProductId:  productObjId,
		ProductImg: "/" + path.Base(tempFile.Name()),
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}

	ser.productRepo.AddProductImg(productImg)

	return nil
}

func (ser *productService) AddProductForSell(productSell dto.CreateProductsSellingDTO) error {

	// check the product is already for sell
	err := ser.CheckProductAlreadyExits(productSell.ProductId)
	if err != nil {
		return err
	}

	_, prodIDVal := helper.ValidatePrimitiveId(productSell.ProductId)
	_, catIdVal := helper.ValidatePrimitiveId(productSell.CategoryId)

	if prodIDVal != nil {
		return prodIDVal
	}

	if catIdVal != nil {
		return catIdVal
	}

	productForSell := new(models.ProductsSelling).SetProductsSelling(productSell)
	ser.productRepo.SellTheProduct(productForSell)
	return nil
}

func (ser *productService) CheckProductAlreadyExits(productId string) error {
	productObjId, prodIDVal := helper.ValidatePrimitiveId(productId)

	if prodIDVal != nil {
		return prodIDVal
	}

	return ser.productRepo.CheckProductAlreadyExits(productObjId)
}

func (ser *productService) ProductsForSelling() (interface{}, error) {
	return ser.productRepo.ProductsForSelling()
}
