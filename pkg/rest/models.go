package rest

type ResponseData struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	TraceID string      `json:"traceid"`
	Data    interface{} `json:"data"`
}

type CreateUserRequestData struct {
	DisplayName string `json:"displayname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	UserName    string `json:"username" binding:"required"`
}

type UpdateUserRequestData struct {
	DisplayName string `json:"displayname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	UserName    string `json:"username" binding:"required"`
}

type CreateCommentRequestData struct {
	UserName string `json:"username" binding:"required"`
	Message  string `json:"message" binding:"required"`
}
