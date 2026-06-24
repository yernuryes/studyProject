package handlers

import (
	"net/http"
	"userService/dto"
	"userService/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	tokenResponse, err := services.Login(
		req.Username,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}

func Refresh(c *gin.Context) {

	var req dto.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	response, err := services.Refresh(req.RefreshToken)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
