package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/repositories"
)

// Potential tech debt -- abstract every controller to share interface, too complicated for now (for me)
type StoryController interface {
	GetRandomStory(c *gin.Context)
	CreateStoryFromURL(c *gin.Context)
	GetStoryById(c *gin.Context)
	DeleteStory(c *gin.Context)
	GetStoriesByAuthor(c *gin.Context)
	GetStoriesByTag(c *gin.Context)
}

type StoryControllerImpl struct {
	StoryRepo repositories.StoryRepository
}

// Main functions
func (sc StoryControllerImpl) GetRandomStory(c *gin.Context) {
	c.JSON(200, "pong")
}

func (sc StoryControllerImpl) CreateStoryFromURL(c *gin.Context) {
}

func (sc StoryControllerImpl) GetStoryById(c *gin.Context) {
	// Story ID validation
	stringId := c.Param("id")
	id, err := validateId(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	storyFromId, err := sc.StoryRepo.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &storyFromId)
}

func (sc StoryControllerImpl) DeleteStory(c *gin.Context) {
	stringId := c.Param("id")
	id, err := validateId(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = sc.StoryRepo.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (sc StoryControllerImpl) GetStoriesByAuthor(c *gin.Context) {

}

func (sc StoryControllerImpl) GetStoriesByTag(c *gin.Context) {
}

// Helper functions
func validateId(stringId string) (uint64, error) {
	id, err := strconv.ParseUint(stringId, 10, 32)
	if err != nil {
		return 0, errors.New(enums.ErrorInvalidStoryId)
	}
	return id, nil
}
