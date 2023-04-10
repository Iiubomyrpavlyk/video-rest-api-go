package dto

type GetAllVideosRequest struct {
	Limit  int `form:"limit,default=25"  binding:"numeric,min=0,max=25"`
	Offset int `form:"offset,default=0" binding:"numeric,min=0"`
}

type CreateVideoDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Duration    int64  `json:"duration" binding:"required"`
	Thumbnail   string `json:"thumbnail"`
	Tags        string `json:"tags"`
	ChannelId   string `json:"channel_id" db:"channel_id" binding:"required"`
}

type UpdateVideoDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Duration    int64  `json:"duration" binding:"required"`
	Thumbnail   string `json:"thumbnail" binding:"required"`
	Tags        string `json:"tags" binding:"required"`
}
