package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-Blog/db"

	"github.com/go-Blog/model"
)

// AddArticle add new article
func AddArticle(w http.ResponseWriter, r *http.Request) {
	article := new(model.Articles)
	json.NewDecoder(r.Body).Decode(&article)
	log.Printf("requet data info :%#v", article)
	if article.UserID == 0 {
		uID := r.Context().Value("uId").(string)
		article.UserID, _ = strconv.Atoi(uID)
	}
	article.ArticleDate = time.Now().Format("2006-01-02 15:04:05")
	if article.ArticleTitle == "" || article.ArticleContent == "" {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "article title or content can not be null",
		})
		return
	}
	stmt, _ := db.DB.Prepare("INSERT INTO article(userId,articleTitle,articleContent,articleViews,articleCommentCount,articleDate,articleLikeCount) Values(?,?,?,?,?,?,?)")
	defer stmt.Close()

	_, err := stmt.Exec(article.UserID, article.ArticleTitle, article.ArticleContent, article.ArticleViews, article.ArticleCommentCount, article.ArticleDate, article.ArticleLikeCount)
	if err != nil {
		log.Printf("create article failed,error:%v\n", err.Error())
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "article create failed",
		})
		return
	}
	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "ok",
	})

}
