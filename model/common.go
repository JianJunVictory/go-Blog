package model

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
type JwtToken struct {
	Token string `json:"token"`
}
