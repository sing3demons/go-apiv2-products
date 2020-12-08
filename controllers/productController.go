package controllers

import (
	"app/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

type createProductForm struct {
	Name     string                `form:"name" binding:"required"`
	Desc     string                `form:"desc" binding:"required"`
	Category string                `form:"category" binding:"required"`
	Price    int64                 `form:"price" binding:"required"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
}

type productResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Image    string `json:"image"`
	Category string `json:"category"`
	Price    int64  `json:"price"`
}

func (p *Product) FindAll(ctx *gin.Context) {
	var products []models.Product
	if err := p.DB.Find(&products).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	
	var serializedArticle []productResponse
	copier.Copy(&serializedArticle, &products)
	ctx.JSON(http.StatusOK, gin.H{"products": serializedArticle})
}

func (p *Product) Create(ctx *gin.Context) {
	var form createProductForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	copier.Copy(&product, &form)

	if err := p.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	p.setProductImage(ctx, &product)
	serializedArticle := productResponse{}
	copier.Copy(&serializedArticle, &product)
	ctx.JSON(http.StatusCreated, gin.H{"product": serializedArticle})
}

func (a *Product) setProductImage(ctx *gin.Context, product *models.Product) error {
	file, err := ctx.FormFile("image")
	if err != nil || file == nil {
		return err
	}

	if product.Image != "" {
		product.Image = strings.Replace(product.Image, os.Getenv("HOST"), "", 1)
		pwd, _ := os.Getwd()
		os.Remove(pwd + product.Image)
	}

	path := "uploads/product/" + strconv.Itoa(int(product.ID))
	os.MkdirAll(path, 0755)
	filename := path + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		return err
	}

	product.Image = os.Getenv("HOST") + "/" + filename
	a.DB.Save(product)

	return nil
}
