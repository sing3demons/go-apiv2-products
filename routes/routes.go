package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

//Serve - init
func Serve(r *gin.Engine) {
	v2 := r.Group("/api/v2")
	productsGroup := v2.Group("products")
	productController := controllers.Product{}

	productsGroup.GET("", productController.FindAll)
}
