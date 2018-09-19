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
	// TODO:insert new article
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

// UpdateArticle update article data
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := new(model.Articles)
	json.NewDecoder(r.Body).Decode(&article)
	log.Printf("requet data info :%#v", article)
	if article.ArticleTitle == "" || article.ArticleContent == "" {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "article title or content can not be null",
		})
		return
	}
	// TODO:update article info
	stmt, _ := db.DB.Prepare(`UPDATE article SET articleTitle=?,articleContent=?,articleViews=? WHERE articleId=?`)
	defer stmt.Close()
	_, err := stmt.Exec(article.ArticleTitle, article.ArticleContent, article.ArticleViews, article.ArticleID)
	if err != nil {
		log.Printf("update article failed,error:%v\n", err.Error())
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "article update failed",
		})
		return
	}
	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "ok",
	})
}

// DeleteArticle delete an article
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	article := new(model.Articles)
	json.NewDecoder(r.Body).Decode(&article)
	log.Printf("requet data info :%#v", article)
	if article.ArticleID == 0 {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "article ID can not be null",
		})
		return
	}
	// TODO:delete an article
	stmt, _ := db.DB.Prepare(`DELETE FROM article WHERE articleId=?`)
	defer stmt.Close()

	_, err := stmt.Exec(article.ArticleID)
	if err != nil {
		log.Printf("delete article failed,error:%v\n", err.Error())
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: " delete an article failed",
		})
		return
	}
	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "ok",
	})
}

// FindArticle find current user article
func FindArticle(w http.ResponseWriter, r *http.Request) {
	uID := r.Context().Value("uId").(string)
	UserID, _ := strconv.Atoi(uID)
	var articles []model.Articles
	article := model.Articles{}
	stmt, _ := db.DB.Prepare(`SELECT * FROM article WHERE userId=?`)
	defer stmt.Close()

	rows, err := stmt.Query(UserID)
	if err != nil {
		log.Println("get article failed")
		log.Println(err)
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: " get article failed",
		})
		return
	}
	for rows.Next() {
		rows.Scan(&article.ArticleID, &article.UserID, &article.ArticleTitle, &article.ArticleContent, &article.ArticleViews, &article.ArticleCommentCount, &article.ArticleDate, &article.ArticleLikeCount)
		articles = append(articles, article)
	}
	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "ok",
		Data:    articles,
	})

}
