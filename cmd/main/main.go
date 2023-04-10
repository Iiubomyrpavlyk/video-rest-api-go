package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	videos "video-rest-api"
	"video-rest-api/internal/adapters/db/postgres"
	v1 "video-rest-api/internal/controller/http/v1"
	"video-rest-api/internal/domain/service"
	"video-rest-api/internal/domain/usecase/channel"
	"video-rest-api/internal/domain/usecase/video"
	"video-rest-api/pkg/client/posgresql"
)

func main() {

	// Setup logger logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while initializing config: %s", err.Error())
	}

	dbConfig := posgresql.DatabaseConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := posgresql.NewClient(dbConfig)

	if err != nil {
		log.Fatalf("error occurred while connecting to database: %s", err)
	}

	defer db.Close()

	r := gin.New()

	channelStorage := postgres.NewChannelStorage(db)
	channelService := service.NewChannelService(channelStorage)

	videoStorage := postgres.NewVideoStorage(db)
	videoService := service.NewVideoService(videoStorage)

	channelUseCase := channel.NewChannelUseCase(channelService)
	channelHandler := v1.NewChannelHandler(channelUseCase)

	channelHandler.Register(r.Group("/api/channels"))

	videoUseCase := video.NewVideoUseCase(videoService, channelService)
	videoHandler := v1.NewVideoHandler(videoUseCase)

	videoHandler.Register(r.Group("/api/videos"))

	srv := new(videos.Server)

	if err := srv.Run(viper.GetString("port"), r); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
