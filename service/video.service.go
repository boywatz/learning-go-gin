package service

import "github/learning/go-gin/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
type videoService struct {
	videos []entity.Video
}

// FindAll implements VideoService.
func (service *videoService) FindAll() []entity.Video {
	if service.videos == nil {
		return []entity.Video{}
	}

	return service.videos
}

// Save implements VideoService.
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func New() VideoService {
	return &videoService{}
}
