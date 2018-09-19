package model

// User struct
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

// Goduser struct
type Goduser struct {
	CUST_ID int    `json:"id"`
	NAME    string `json:"name"`
	AGE     int    `json:"age"`
}
