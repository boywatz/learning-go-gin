package main

import (
	"github/learning/go-gin/controller"
	middleware "github/learning/go-gin/middlewares"
	"github/learning/go-gin/service"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	file, err := os.Create("server.log")
	if err != nil {
		log.Printf("Error creating file: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		video, err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(201, video)
	})

	server.Run(":8080")
}
