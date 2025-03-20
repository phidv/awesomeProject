package request

type GetUserRequest struct {
	Id string `json:"id" binding:"required,uuid" form:"id"`
}

type ListUsersRequest struct {
}
