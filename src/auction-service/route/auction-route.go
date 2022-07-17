package route

import (
	"chilindo/src/auction-service/controller"
	"github.com/gin-gonic/gin"
)

type IAuctionRoute interface {
	GetRouter()
}

type AuctionRoute struct {
	AuctionController controller.IAuctionController
	Router            *gin.Engine
}

func NewAuctionRoute(auctionController controller.IAuctionController, router *gin.Engine) *AuctionRoute {
	return &AuctionRoute{AuctionController: auctionController, Router: router}
}

func (a AuctionRoute) GetRouter() {
	auctionRoute := a.Router.Group("/chilindo/auction/")
	{
		auctionRoute.POST("/create", a.AuctionController.CreateAuction)
	}
}
