package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-Blog/model"
)

func TestRegister(t *testing.T) {
	user := model.User{Email: "enjoyass@outlook.com", Password: "12345678"}
	postData, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(postData))
	if err != nil {
		log.Println(err)
	}
	// req.Header.Set("Conetent-Type", "application/json")
	rsp := httptest.NewRecorder()
	Register(rsp, req)

}
