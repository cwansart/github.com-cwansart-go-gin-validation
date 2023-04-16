package config

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, CurrentConfig)
}

func UpdateConfig(c *gin.Context) {
	var newConfig Config
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		fmt.Println("an error occured:", err)
		return
	}

	// Manual validation, requires validation as "validate" tag on struct.
	// v := validator.New()
	// if err := v.Struct(&newConfig); err != nil {
	// 	fmt.Println("validation for config failed:", err)
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }

	CurrentConfig = newConfig

	c.JSON(http.StatusOK, newConfig)
}
