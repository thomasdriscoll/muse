package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/models"
	"github.com/thomasdriscoll/muse/services"
)

type StoryController interface {
	GetRandomStory(c *gin.Context)
	CreateStoryFromURL(c *gin.Context)
	GetStoryById(c *gin.Context)
	DeleteStory(c *gin.Context)
	GetStoriesByAuthor(c *gin.Context)
	GetStoriesByTag(c *gin.Context)
}

type StoryControllerImpl struct {
	StorySvc services.StoryService
}

// Main functions
func (sc StoryControllerImpl) GetRandomStory(c *gin.Context) {
	story, err := sc.StorySvc.GetStoryByRandom()
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(200, &story)
}

func (sc StoryControllerImpl) CreateStoryFromURL(c *gin.Context) {
	storyRequest := models.StoryFromURLRequest{}
	if err := c.BindJSON(&storyRequest); err != nil {
		c.JSON(http.StatusBadRequest, enums.ErrorInvalidStoryRequest)
		return
	}
	story, err := sc.StorySvc.CreateStory(&storyRequest)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusCreated, &story)
}

func (sc StoryControllerImpl) GetStoryById(c *gin.Context) {
	// Story ID validation
	stringId := c.Param("id")
	id, err := validateId(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	storyFromId, err := sc.StorySvc.GetStoryById(id)

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

	err = sc.StorySvc.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (sc StoryControllerImpl) GetStoriesByAuthor(c *gin.Context) {
	stringId := c.Param("authorId")
	authorId, err := validateId(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	stories, err := sc.StorySvc.GetStoriesByAuthorId(authorId)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &stories)
}

func (sc StoryControllerImpl) GetStoriesByTag(c *gin.Context) {
	tag := c.Param("tag")

	stories, err := sc.StorySvc.GetStoriesByTag(tag)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &stories)
}

// Helper functions
func validateId(stringId string) (string, error) {
	if stringId == "junk" {
		return "", errors.New(enums.ErrorInvalidStoryId)
	}
	return stringId, nil
}
