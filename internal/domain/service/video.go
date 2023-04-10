package service

import (
	"time"
	"video-rest-api/internal/controller/http/dto"
	"video-rest-api/internal/domain/entity"
)

type VideoStorage interface {
	Create(video entity.Video) (string, error)
	GetById(videoId string) (entity.Video, error)
	GetAll(q string, limit, offset int) ([]entity.Video, error)
	Delete(videoId string) error
	Update(video entity.Video) error
	GetAllByChannelId(channelId string) ([]entity.Video, error)
}

type videoService struct {
	storage VideoStorage
}

func NewVideoService(storage VideoStorage) *videoService {
	return &videoService{storage: storage}
}

func (s *videoService) CreateVideo(dto dto.CreateVideoDTO) (string, error) {

	v := entity.Video{
		Id:          "",
		Title:       dto.Title,
		Description: dto.Description,
		Duration:    dto.Duration,
		Thumbnail:   dto.Thumbnail,
		Tags:        dto.Tags,
		PublishedAt: time.Now(),
		ChannelId:   dto.ChannelId,
	}

	return s.storage.Create(v)
}

func (s *videoService) GetVideoById(videoId string) (entity.Video, error) {
	return s.storage.GetById(videoId)
}

func (s *videoService) GetAllVideos(q string, limit, offset int) ([]entity.Video, error) {
	return s.storage.GetAll(q, limit, offset)
}

func (s *videoService) DeleteVideo(videoId string) error {
	return s.storage.Delete(videoId)
}

func (s *videoService) UpdateVideo(videoId string, videoDTO dto.UpdateVideoDTO) error {

	video, err := s.storage.GetById(videoId)
	if err != nil {
		return err
	}

	return s.storage.Update(entity.Video{
		Id:          videoId,
		Title:       videoDTO.Title,
		Description: videoDTO.Description,
		Duration:    videoDTO.Duration,
		Thumbnail:   videoDTO.Thumbnail,
		Tags:        videoDTO.Tags,
		PublishedAt: video.PublishedAt,
		ChannelId:   video.ChannelId,
	})
}

func (s *videoService) GetAllByChannelId(channelId string) ([]entity.Video, error) {
	return s.storage.GetAllByChannelId(channelId)
}
