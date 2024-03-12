// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserType int64  `json:"userType"`
	Avatar   string `json:"avatar"`
	Status   int64  `json:"status"`
}

type LoginRequest struct {
	UserKey  string `json:"userKey"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    bool   `json:"data"`
}