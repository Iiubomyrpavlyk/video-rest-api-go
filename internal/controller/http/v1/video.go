package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	h2 "video-rest-api/___pkg/handler"
	"video-rest-api/internal/controller/http/dto"
	"video-rest-api/internal/domain/entity"
)

type VideoUseCase interface {
	CreateVideo(dto dto.CreateVideoDTO) (string, error)
	GetAllVideos(q string, limit, offset int) ([]entity.Video, error)
	GetVideoById(videoId string) (entity.Video, error)
	DeleteVideo(videoId string) error
	UpdateVideo(videoId string, dto dto.UpdateVideoDTO) error
}

type videoHandler struct {
	videoUseCase VideoUseCase
}

func NewVideoHandler(useCase VideoUseCase) *videoHandler {
	return &videoHandler{videoUseCase: useCase}
}

func (h *videoHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.createVideo)
	router.GET("/", h.getAllVideos)
	router.GET("/:id", h.getVideoById)
	router.DELETE("/:id", h.deleteVideo)
	router.PUT("/:id", h.updateVideo)
}

func (h *videoHandler) createVideo(ctx *gin.Context) {
	var requestBody dto.CreateVideoDTO

	// Abort with HTTP Status Code 400 if channel request body is not valid
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, h2.ResponseError{Message: err.Error()})
		return
	}

	id, err := h.videoUseCase.CreateVideo(requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	// Pass HTTP Status Code 201
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *videoHandler) getAllVideos(ctx *gin.Context) {
	q := ctx.Query("q")

	if len(q) == 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{})
		return
	}

	var request dto.GetAllVideosRequest

	// Validate query params
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, h2.ResponseError{Message: err.Error()})
		return
	}

	videos, err := h.videoUseCase.GetAllVideos(q, request.Limit, request.Offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, videos)
}

func (h *videoHandler) getVideoById(ctx *gin.Context) {
	videoId := ctx.Param("id")

	channel, err := h.videoUseCase.GetVideoById(videoId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channel)
}

func (h *videoHandler) deleteVideo(ctx *gin.Context) {
	videoId := ctx.Param("id")

	if err := h.videoUseCase.DeleteVideo(videoId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
	ctx.Abort()
}

func (h *videoHandler) updateVideo(ctx *gin.Context) {
	videoId := ctx.Param("id")
	var requestBody dto.UpdateVideoDTO

	// Abort with HTTP Status Code 400 if channel request body is not valid
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, h2.ResponseError{Message: err.Error()})
		return
	}

	if err := h.videoUseCase.UpdateVideo(videoId, requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
	ctx.Abort()
}
