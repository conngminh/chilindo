package controller

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"chilindo/src/product-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	imageId = "imageId"
)

type ProductImageController interface {
	CreateImage(c *gin.Context)
	GetImage(c *gin.Context)
	GetImageByID(c *gin.Context)
	DeleteImage(c *gin.Context)
}

func (p productImageController) GetImageByID(c *gin.Context) {
	//TODO implement me
	var dto dto.ImageDTO
	dto.ImageId = c.Param(imageId)
	c.Set(imageId, dto.ImageId)
	image, err := p.productImageService.GetImageByID(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error get option",
		})
		log.Println("GetOptionById: Error call service in pkg controller", err)
		c.Abort()
		return
	}
	if image == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Option not found",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, image)
}

func (p productImageController) DeleteImage(c *gin.Context) {
	//TODO implement me
	oId := c.Param(imageId)
	var dto dto.ImageDTO
	dto.ImageId = oId
	image, err := p.productImageService.DeleteImage(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to delete option",
		})
		log.Println("DeleteOption: Error to parse oId", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, image)
}

func (p productImageController) GetImage(c *gin.Context) {
	//TODO implement me
	id := c.Param(productId)
	var dto dto.ProductIdDTO
	dto.ProductId = id
	images, err := p.productImageService.GetImage(&dto)
	if err != nil {
		log.Println("GetOptions: error in controller package", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to get options",
		})
		c.Abort()
		return
	}
	if images == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Not found options",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, images)
}

func (p productImageController) CreateImage(c *gin.Context) {
	//TODO implement me
	var imageBody *entity.ProductImages
	if err := c.ShouldBindJSON(&imageBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create option",
		})
		log.Println("CreateOption: Error to ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	dtoImage := dto.NewCreateImageDTO(imageBody)
	dtoImage.Image.ProductId = c.Param(productId)
	image, err := p.productImageService.CreateImage(dtoImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create option",
		})
		log.Println("CreateOption: Error to CreateOption in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, image)
}

type productImageController struct {
	productImageService service.ProductImageService
}

func NewProductImageController(productImageService service.ProductImageService) *productImageController {
	return &productImageController{productImageService: productImageService}
}
