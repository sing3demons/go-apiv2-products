package routes

import (
	"app/config"
	"app/controllers"

	"github.com/gin-gonic/gin"
)

//Serve - init
func Serve(r *gin.Engine) {
	db := config.GetDB()
	v2 := r.Group("/api/v2")
	productsGroup := v2.Group("products")
	productController := controllers.Product{DB: db}

	productsGroup.GET("", productController.FindAll)
	productsGroup.POST("", productController.Create)
}
