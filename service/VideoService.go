package service

import "github.com/EDDYCJY/go-gin-example/entity"

type IVideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type VideoService struct {
	videos []entity.Video
}

func New() IVideoService {
	return &VideoService{}
}

func (service *VideoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *VideoService) FindAll() []entity.Video {
	return service.videos
}
