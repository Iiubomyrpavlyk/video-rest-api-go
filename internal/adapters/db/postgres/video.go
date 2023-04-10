package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"video-rest-api/internal/domain/entity"
	"video-rest-api/pkg/client/posgresql"
)

type videoStorage struct {
	db *sqlx.DB
}

func NewVideoStorage(db *sqlx.DB) *videoStorage {
	return &videoStorage{db: db}
}

func (v *videoStorage) Create(video entity.Video) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description, duration, thumbnail, tags, published_at, channel_id) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", posgresql.VideosTable)

	row := v.db.QueryRow(query, video.Title, video.Description, video.Duration, video.Thumbnail, video.Tags, video.PublishedAt, video.ChannelId)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (v *videoStorage) GetById(videoId string) (entity.Video, error) {
	query := fmt.Sprintf("SELECT id, title, description, duration, thumbnail, tags, published_at, channel_id FROM %s WHERE id = '%s'", posgresql.VideosTable, videoId)

	var video entity.Video

	if err := v.db.Get(&video, query); err != nil {
		return video, err
	}

	return video, nil
}

func (v *videoStorage) GetAll(q string, limit, offset int) ([]entity.Video, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE title LIKE '%%%s%%' LIMIT %d OFFSET %d", posgresql.VideosTable, q, limit, offset)

	videos := make([]entity.Video, 0)

	err := v.db.Select(&videos, query)

	if err != nil {
		// Handle no rows error
		if err == sql.ErrNoRows {
			return videos, nil
		}

		return videos, err
	}

	return videos, nil
}

func (v *videoStorage) Delete(videoId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", posgresql.VideosTable, videoId)

	if _, err := v.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (v *videoStorage) Update(video entity.Video) error {
	query := fmt.Sprintf("UPDATE %s SET title = '%s', description = '%s', duration = '%d', thumbnail = '%s', tags = '%s' WHERE id = '%s'", posgresql.VideosTable, video.Title, video.Description, video.Duration, video.Thumbnail, video.Tags, video.Id)

	if _, err := v.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (v *videoStorage) GetByChannelId(channelId string) ([]entity.Video, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE channel_Id = '%s'", posgresql.VideosTable, channelId)

	videos := make([]entity.Video, 0)

	err := v.db.Select(&videos, query)

	if err != nil {
		// Handle no rows error
		if err == sql.ErrNoRows {
			return videos, nil
		}

		return videos, err
	}

	return videos, nil
}
