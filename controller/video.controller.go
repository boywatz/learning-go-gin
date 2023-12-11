package controller

import (
	"github/learning/go-gin/entity"
	"github/learning/go-gin/service"
	"github/learning/go-gin/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) (entity.Video, error)
}

type controller struct {
	videoService service.VideoService
}

// FindAll implements VideoController.
func (c *controller) FindAll() []entity.Video {
	return c.videoService.FindAll()
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return video, err
	}

	err = validate.Struct(video)
	if err != nil {
		return video, err
	}

	c.videoService.Save(video)
	return video, nil
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("warning-text", validators.IsWarningText)

	return &controller{
		videoService: service,
	}
}
