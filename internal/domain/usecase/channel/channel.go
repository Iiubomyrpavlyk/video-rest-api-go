package channel

import (
	"video-rest-api/internal/controller/http/dto"
	"video-rest-api/internal/domain/entity"
)

type Service interface {
	CreateChannel(channel dto.CreateChannelDTO) (string, error)
	GetChannelById(channelId string) (entity.Channel, error)
	GetAllChannels(q string, limit, offset int) ([]entity.Channel, error)
	DeleteChannel(channelId string) error
	UpdateChannel(channelId string, channelDTO dto.UpdateChannelDTO) error
}

type VideoService interface {
	GetAllByChannelId(channelId string) ([]entity.Video, error)
}

type channelUseCase struct {
	channelService Service
	videoService   VideoService
}

func NewChannelUseCase(channelService Service, service VideoService) *channelUseCase {
	return &channelUseCase{channelService: channelService, videoService: service}
}

func (u channelUseCase) CreateChannel(dto dto.CreateChannelDTO) (string, error) {
	return u.channelService.CreateChannel(dto)
}

func (u channelUseCase) GetAllChannels(q string, limit, offset int) ([]entity.Channel, error) {
	return u.channelService.GetAllChannels(q, limit, offset)
}

func (u channelUseCase) GetChannelById(id string) (entity.Channel, error) {
	return u.channelService.GetChannelById(id)
}

func (u channelUseCase) UpdateChannel(id string, channelDTO dto.UpdateChannelDTO) error {
	return u.channelService.UpdateChannel(id, channelDTO)
}

func (u channelUseCase) DeleteChannel(id string) error {
	return u.channelService.DeleteChannel(id)
}

func (u channelUseCase) GetAllByChannelId(id string) ([]entity.Video, error) {
	return u.videoService.GetAllByChannelId(id)
}
