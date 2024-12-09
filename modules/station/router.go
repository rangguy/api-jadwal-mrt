package station

import (
	"github.com/gin-gonic/gin"
	"github.com/rangguy/api-jadwal-mrt.git/common/response"
	"net/http"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)

	})

	station.GET("/:id", func(c *gin.Context) {
		CheckScheduleByStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStations()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	// return response
	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Success",
		Data:    datas,
	})
}

func CheckScheduleByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	datas, err := service.CheckScheduleByStation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})

		return
	}

	// return response
	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Success",
		Data:    datas,
	})
}
