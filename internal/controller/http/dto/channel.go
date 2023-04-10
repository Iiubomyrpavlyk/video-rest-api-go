package dto

type GetAllChannelsRequest struct {
	Limit  int `form:"limit,default=25"  binding:"numeric,min=0,max=25"`
	Offset int `form:"offset,default=0" binding:"numeric,min=0"`
}

type CreateChannelDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Thumbnail   string `json:"thumbnail"`
}

type UpdateChannelDTO struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Thumbnail       string `json:"thumbnail"`
	SubscriberCount int    `json:"subscriber_count" db:"subscriber_count"`
}
