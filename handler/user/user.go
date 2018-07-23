package user

import "gingob/model"

// LoginRequest is the login request params struct
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateRequest is the create user request params struct
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateResponse is the create user request params struct
type CreateResponse struct {
	Username string `json:"username"`
}

// ListRequest is the user list request struct
type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

// ListResponse is the use list response struct
type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
