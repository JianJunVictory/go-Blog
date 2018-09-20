package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-Blog/model"
)

// AddComment add comment to article
func AddComment(w http.ResponseWriter, r *http.Request) {
	comment := new(model.Comments)
	json.NewDecoder(r.Body).Decode(comment)
	log.Printf("request comment data info %#v ", comment)

	uID := r.Context().Value("uId").(string)
	UserID, _ := strconv.Atoi(uID)
	comment.UserID = UserID
	w.Write([]byte("aaaaaaaaaaaaaaaaaaaa"))
}
