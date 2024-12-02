package station

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/station")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)

	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStations()
	if err != nil {
		// handle error
	}

	// return response
	c.JSON(http.StatusOK, datas)
}
