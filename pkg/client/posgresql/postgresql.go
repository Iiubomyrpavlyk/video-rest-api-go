package posgresql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	ChannelsTable       = "channel"
	VideosTable         = "video"
	PlaylistsTable      = "playlist"
	PlaylistVideosTable = "playlist_videos"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewClient(config DatabaseConfig) (*sqlx.DB, error) {
	dbClient, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))

	if err != nil {
		return nil, err
	}

	return dbClient, nil
}
