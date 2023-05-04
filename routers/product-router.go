package routers

import (
	"github.com/aniket0951.com/product-service/controller"
	"github.com/aniket0951.com/product-service/repository"
	"github.com/aniket0951.com/product-service/services"
	"github.com/gin-gonic/gin"
)

var (
	prodrepo       = repository.NewProductRepository()
	prodservice    = services.NewProductService(prodrepo)
	prodcontroller = controller.NewProductController(prodservice)
)

func ProductRouter(router *gin.Engine) {
	productRoutes := router.Group("/api/product")
	{
		productRoutes.POST("/create-product", prodcontroller.CreateProduct)
		productRoutes.PUT("/update-product", prodcontroller.UpdateProduct)
		productRoutes.DELETE("/remove-product", prodcontroller.DeleteProduct)
		productRoutes.GET("/seller-products", prodcontroller.ProductsBySeller)
		productRoutes.PUT("/increase-decrease-product", prodcontroller.IncreaseDecreaseProduct)
		productRoutes.POST("/add-product-img", prodcontroller.AddProductImage)
	}

	productSellRoutes := router.Group("/api/product/sell")
	{
		productSellRoutes.POST("/add-product-sell", prodcontroller.AddProductForSell)
		productSellRoutes.GET("/products-for-selling", prodcontroller.ProductsForSelling)
	}

}
