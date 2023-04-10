package video

import (
	"errors"
	"video-rest-api/internal/controller/http/dto"
	"video-rest-api/internal/domain/entity"
	"video-rest-api/internal/domain/usecase/channel"
)

type Service interface {
	CreateVideo(videoDTO dto.CreateVideoDTO) (string, error)
	GetVideoById(videoId string) (entity.Video, error)
	GetAllVideos(q string, limit, offset int) ([]entity.Video, error)
	DeleteVideo(videoId string) error
	UpdateVideo(videoId string, videoDTO dto.UpdateVideoDTO) error
	GetByChannelId(channelId string) ([]entity.Video, error)
}

type videoUseCase struct {
	videoService   Service
	channelService channel.Service
}

func NewVideoUseCase(videoService Service, channelService channel.Service) *videoUseCase {
	return &videoUseCase{
		videoService:   videoService,
		channelService: channelService,
	}
}

func (u videoUseCase) CreateVideo(dto dto.CreateVideoDTO) (string, error) {
	if _, err := u.channelService.GetChannelById(dto.ChannelId); err != nil {
		return "", errors.New("unknown channel id error")
	}

	return u.videoService.CreateVideo(dto)
}

func (u videoUseCase) GetAllVideos(q string, limit, offset int) ([]entity.Video, error) {
	return u.videoService.GetAllVideos(q, limit, offset)
}

func (u videoUseCase) GetVideoById(id string) (entity.Video, error) {
	return u.videoService.GetVideoById(id)
}

func (u videoUseCase) UpdateVideo(id string, videoDTO dto.UpdateVideoDTO) error {
	return u.videoService.UpdateVideo(id, videoDTO)
}

func (u videoUseCase) DeleteVideo(id string) error {
	return u.videoService.DeleteVideo(id)
}

func (u videoUseCase) GetByChannelId(id string) ([]entity.Video, error) {
	return u.videoService.GetByChannelId(id)
}
