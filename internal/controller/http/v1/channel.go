package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	h2 "video-rest-api/___pkg/handler"
	"video-rest-api/internal/controller/http/dto"
	"video-rest-api/internal/domain/entity"
)

const (
	channelURL  = "/channels/:id"
	channelsURL = "/channels"
)

type ChannelUseCase interface {
	CreateChannel(dto dto.CreateChannelDTO) (string, error)
	GetAllChannels(q string, limit, offset int) ([]entity.Channel, error)
	GetChannelById(id string) (entity.Channel, error)
	UpdateChannel(id string, channelDTO dto.UpdateChannelDTO) error
	DeleteChannel(id string) error
}

type channelHandler struct {
	channelUseCase ChannelUseCase
}

func NewChannelHandler(useCase ChannelUseCase) *channelHandler {
	return &channelHandler{channelUseCase: useCase}
}

func (h *channelHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.createChannel)
	router.GET("/", h.getAllChannels)
	router.GET("/:id", h.getChannelById)
	router.DELETE("/:id", h.deleteChannel)
	router.PUT("/:id", h.updateChannel)
}

func (h *channelHandler) createChannel(ctx *gin.Context) {
	var requestBody dto.CreateChannelDTO

	// Abort with HTTP Status Code 400 if channel request body is not valid
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, h2.ResponseError{Message: err.Error()})
		return
	}

	id, err := h.channelUseCase.CreateChannel(requestBody)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	// Pass HTTP Status Code 201
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *channelHandler) getAllChannels(ctx *gin.Context) {
	q := ctx.Query("q")

	if len(q) == 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{})
		return
	}

	var request dto.GetAllChannelsRequest

	// Validate query params
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, h2.ResponseError{Message: err.Error()})
		return
	}

	channels, err := h.channelUseCase.GetAllChannels(q, request.Limit, request.Offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channels)
}

func (h *channelHandler) getChannelById(ctx *gin.Context) {
	channelId := ctx.Param("id")

	channel, err := h.channelUseCase.GetChannelById(channelId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channel)
}

func (h *channelHandler) deleteChannel(ctx *gin.Context) {
	channelId := ctx.Param("id")

	if err := h.channelUseCase.DeleteChannel(channelId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
	ctx.Abort()
}

func (h *channelHandler) updateChannel(ctx *gin.Context) {
	channelId := ctx.Param("id")
	var requestBody dto.UpdateChannelDTO

	// Abort with HTTP Status Code 400 if channel request body is not valid
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, h2.ResponseError{Message: err.Error()})
		return
	}

	if err := h.channelUseCase.UpdateChannel(channelId, requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, h2.ResponseError{Message: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
	ctx.Abort()
}
