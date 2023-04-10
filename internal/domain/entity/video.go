package entity

import (
	"time"
)

type Video struct {
	Id          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Duration    int64     `json:"duration" binding:"required"`
	Thumbnail   string    `json:"thumbnail"`
	Tags        string    `json:"tags"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	ChannelId   string    `json:"channel_id" db:"channel_id" binding:"required"`
}
