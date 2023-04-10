package service

import (
	"errors"
	"time"
	"video-rest-api/internal/controller/http/dto"
	"video-rest-api/internal/domain/entity"
)

type ChannelStorage interface {
	Create(channel entity.Channel) (string, error)
	GetById(channelId string) (entity.Channel, error)
	GetAll(q string, limit, offset int) ([]entity.Channel, error)
	Delete(channelId string) error
	Update(channel entity.Channel) error
}

type channelService struct {
	storage ChannelStorage
}

func NewChannelService(storage ChannelStorage) *channelService {
	return &channelService{storage: storage}
}

func (s *channelService) CreateChannel(channel dto.CreateChannelDTO) (string, error) {

	println(channel.Title)

	if chs, _ := s.storage.GetAll(channel.Title, 1, 0); len(chs) > 0 {
		return "", errors.New("channel with the same title already exists")
	}

	ch := entity.Channel{
		Id:              "",
		Title:           channel.Title,
		Description:     channel.Description,
		Thumbnail:       channel.Thumbnail,
		PublishedAt:     time.Now(),
		SubscriberCount: 0,
	}

	return s.storage.Create(ch)
}

func (s *channelService) GetChannelById(channelId string) (entity.Channel, error) {
	return s.storage.GetById(channelId)
}

func (s *channelService) GetAllChannels(q string, limit, offset int) ([]entity.Channel, error) {
	return s.storage.GetAll(q, limit, offset)
}

func (s *channelService) DeleteChannel(channelId string) error {
	return s.storage.Delete(channelId)
}

func (s *channelService) UpdateChannel(channelId string, channelDTO dto.UpdateChannelDTO) error {

	channel, err := s.storage.GetById(channelId)
	if err != nil {
		return err
	}

	return s.storage.Update(entity.Channel{
		Id:              channelId,
		Title:           channelDTO.Title,
		Description:     channelDTO.Description,
		Thumbnail:       channelDTO.Thumbnail,
		PublishedAt:     channel.PublishedAt,
		SubscriberCount: channelDTO.SubscriberCount,
	})
}
