package entity

import "time"

type Playlist struct {
	Id          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Thumbnail   string    `json:"thumbnail"`
	PublishedAt time.Time `json:"published_at"`
	ChannelId   string    `json:"channel_id" binding:"required"`
	Videos      []struct {
		Id string `json:"id"`
	} `json:"videos" binding:"required" db:"-"` // ignore this field in database
}
