package model

type GetUserResponse struct {
	Id      string `json:"id"`
	IsAdmin bool   `json:"usAdmin"`
}
