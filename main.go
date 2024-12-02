package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rangguy/api-jadwal-mrt.git/modules/station"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	station.Initiate(api)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
