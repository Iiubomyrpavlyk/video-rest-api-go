package entity

import (
	"time"
)

type Channel struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Thumbnail       string    `json:"thumbnail"`
	PublishedAt     time.Time `json:"published_at" db:"published_at"`
	SubscriberCount int       `json:"subscriber_count" db:"subscriber_count"`
}
