package main

import (
	"github.com/cwansart/go-gin-validation/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/config", config.GetConfig)
	router.PATCH("/config", config.UpdateConfig)

	router.Run("localhost:2222")
}
