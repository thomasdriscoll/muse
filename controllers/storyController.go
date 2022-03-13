package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/repositories"
	"github.com/thomasdriscoll/muse/services"
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
	StoryRepo     repositories.StoryRepository
	StoryScrapper services.StoryScrapper
}

// Main functions
func (sc StoryControllerImpl) GetRandomStory(c *gin.Context) {
	story, err := sc.StoryRepo.GetStoryByRandom()
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(200, &story)
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

	storyFromId, err := sc.StoryRepo.GetStoryById(id)

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
	stringId := c.Param("authorId")
	authorId, err := validateId(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	stories, err := sc.StoryRepo.GetStoriesByAuthorId(authorId)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &stories)
}

func (sc StoryControllerImpl) GetStoriesByTag(c *gin.Context) {
	tag := c.Param("tag")

	stories, err := sc.StoryRepo.GetStoriesByTag(tag)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &stories)
}

// Helper functions
func validateId(stringId string) (uint64, error) {
	id, err := strconv.ParseUint(stringId, 10, 32)
	if err != nil {
		return 0, errors.New(enums.ErrorInvalidStoryId)
	}
	return id, nil
}
