package controller

import (
	"chilindo/src/auction-service/entity"
	"chilindo/src/auction-service/service"
	"chilindo/src/pkg/pb/product"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAuctionController interface {
	CreateAuction(c *gin.Context)
}

type AuctionController struct {
	AuctionService service.IAuctionService
	ProductClient  product.ProductServiceClient
}

func NewAuctionController(auctionService service.IAuctionService, productClient product.ProductServiceClient) *AuctionController {
	return &AuctionController{AuctionService: auctionService, ProductClient: productClient}
}

func (a AuctionController) CreateAuction(c *gin.Context) {
	var auctionBody *entity.Auction
	if err := c.ShouldBindJSON(&auctionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create auction",
		})
		log.Println("Error to ShouldBindJSON controller", err)
		c.Abort()
		return
	}
	in := &product.GetProductRequest{ProductId: auctionBody.ProductId}
	res, errRes := a.ProductClient.GetProduct(c, in)
	if errRes != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create Auction",
		})
		log.Println("CreateAuction: Error to call productService rpc server", errRes)
		c.Abort()
		return
	}

	if res.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found product",
		})
		log.Println("CreateAuction: product not found")
		c.Abort()
		return
	}

	auctionBody.ProductId = res.GetId()

	createdAuction, errCreateAuction := a.AuctionService.CreateAuction(auctionBody)
	if errCreateAuction != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errCreateAuction.Error(),
		})
		log.Println("CreateAuction: Error create new auction in package controller")
		return
	}
	c.JSON(http.StatusOK, createdAuction)
}
