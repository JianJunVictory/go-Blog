package model

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

type Goduser struct {
	CUST_ID int    `json:"id"`
	NAME    string `json:"name"`
	AGE     int    `json:"age"`
}
