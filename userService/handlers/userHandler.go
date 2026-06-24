package handlers

import (
	"net/http"
	"userService/dto"
	"userService/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := services.CreateUser(
		req.Username,
		req.Email,
		req.LastName,
		req.FirstName,
		req.Password,
		req.Role,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func UpdateUser(c *gin.Context) {

	username := c.Param("username")

	var req dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := services.UpdateUser(
		username,
		req.Email,
		req.FirstName,
		req.LastName,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
}

func ChangePassword(c *gin.Context) {

	username := c.Param("username")

	var req dto.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	err := services.ChangePassword(
		username,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password changed",
	})
}
