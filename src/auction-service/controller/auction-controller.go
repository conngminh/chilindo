package controller

import (
	"chilindo/src/auction-service/entity"
	"chilindo/src/auction-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAuctionController interface {
	CreateAuction(c *gin.Context)
}

type AuctionController struct {
	AuctionService service.IAuctionService
}

func NewAuctionController(auctionService service.IAuctionService) *AuctionController {
	return &AuctionController{AuctionService: auctionService}
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
