package routers

import (
	"github.com/aniket0951.com/product-service/controller"
	"github.com/aniket0951.com/product-service/repository"
	"github.com/aniket0951.com/product-service/services"
	"github.com/gin-gonic/gin"
)

var (
	catrepo       = repository.NewCategoryRepository()
	catservice    = services.NewCategoryService(catrepo)
	catcontroller = controller.NewCategoryController(catservice)
)

func CategoryRouter(router *gin.Engine) {
	catRoutes := router.Group("/api/category")
	{
		catRoutes.POST("/create-category", catcontroller.CreateCategory)
		catRoutes.PUT("/update-category", catcontroller.UpdateCategory)
		catRoutes.DELETE("/remove-category", catcontroller.DeleteCategory)
		catRoutes.GET("/all-categories", catcontroller.GetAllCategory)
		catRoutes.GET("/fetch-category", catcontroller.CategoryById)
	}

}
