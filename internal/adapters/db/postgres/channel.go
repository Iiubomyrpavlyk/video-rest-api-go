package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"video-rest-api/internal/domain/entity"
	"video-rest-api/pkg/client/posgresql"
)

type channelStorage struct {
	db *sqlx.DB
}

func NewChannelStorage(db *sqlx.DB) *channelStorage {
	return &channelStorage{db: db}
}

func (c *channelStorage) Create(channel entity.Channel) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description, thumbnail, published_at, subscriber_count) values ($1, $2, $3, $4, $5) RETURNING id", posgresql.ChannelsTable)

	row := c.db.QueryRow(query, channel.Title, channel.Description, channel.Thumbnail, channel.PublishedAt, channel.SubscriberCount)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *channelStorage) GetById(channelId string) (entity.Channel, error) {
	query := fmt.Sprintf("SELECT id, title, description, thumbnail, published_at, subscriber_count FROM %s WHERE id = '%s'", posgresql.ChannelsTable, channelId)

	var channel entity.Channel

	if err := c.db.Get(&channel, query); err != nil {
		log.Printf("ChannelRepository error: %s", err.Error())
		return channel, err
	}

	return channel, nil
}

func (c *channelStorage) GetAll(q string, limit, offset int) ([]entity.Channel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE title LIKE '%%%s%%' LIMIT %d OFFSET %d", posgresql.ChannelsTable, q, limit, offset)

	channel := make([]entity.Channel, 0)

	err := c.db.Select(&channel, query)

	if err != nil {
		log.Printf("ChannelRepository error: %s", err.Error())
		// Handle no rows error
		if err == sql.ErrNoRows {
			return channel, nil
		}

		return channel, err
	}

	return channel, nil
}

func (c *channelStorage) Delete(channelId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", posgresql.ChannelsTable, channelId)

	if _, err := c.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (c *channelStorage) Update(channel entity.Channel) error {

	query := fmt.Sprintf("UPDATE %s SET title = '%s', description = '%s', thumbnail = '%s', subscriber_count = %d WHERE id = '%s'", posgresql.ChannelsTable, channel.Title, channel.Description, channel.Thumbnail, channel.SubscriberCount, channel.Id)

	if _, err := c.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (c *channelStorage) GetByTitle(title string) (entity.Channel, error) {
	query := fmt.Sprintf("SELECT id, title, description, thumbnail, published_at, subscriber_count FROM %s WHERE title = %s", posgresql.ChannelsTable, title)

	var channel entity.Channel

	if err := c.db.Get(&channel, query); err != nil {
		return channel, err
	}

	return channel, nil
}
