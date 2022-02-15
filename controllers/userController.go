package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	GetSavedStoriesByUser(c *gin.Context)
}

type UserControllerImpl struct{}

func (uc UserControllerImpl) GetUser(c *gin.Context) {
}

func (uc UserControllerImpl) CreateUser(c *gin.Context) {
}

func (uc UserControllerImpl) GetSavedStoriesByUser(c *gin.Context) {
}
