package controller

import (
	"github/learning/go-gin/entity"
	"github/learning/go-gin/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	videoService service.VideoService
}

// FindAll implements VideoController.
func (c *controller) FindAll() []entity.Video {
	return c.videoService.FindAll()
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.videoService.Save(video)
	return video
}

func New(service service.VideoService) VideoController {
	return &controller{
		videoService: service,
	}
}
