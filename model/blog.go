package model

// Articles struct
type Articles struct {
	ArticleID           int    `json:"articleId"`
	UserID              int    `json:"userId"`
	ArticleTitle        string `json:"articleTitle"`
	ArticleContent      string `json:"articleContent"`
	ArticleViews        string `json:"articleViews"`
	ArticleCommentCount int    `json:"articleCommentCount"`
	ArticleDate         string `json:"articleDate"`
	ArticleLikeCount    int    `json:"articleLikeCount"`
}

// Comments struct
type Comments struct {
	CommentID        int    `json:"commentId"`
	UserID           int    `json:"userId"`
	ArticleID        int    `json:"articleId"`
	CommentLikeCount int    `json:"commentLikeCount"`
	CommentDate      string `json:"commentDate"`
	CommentContent   string `json:"commentContent"`
	ParentCommentID  int    `json:"parentCommentId"`
}

// Sorts struct
type Sorts struct {
	SortID          int    `json:"srotId"`
	SortName        string `json:"sortName"`
	SortAlias       string `json:"sortAlias"`
	SortDescription string `json:"sortDescription"`
	ParentSortID    int    `json:"parentSortId"`
}

// ArticleSort struct
type ArticleSort struct {
	ArticleID int `json:"articleId"`
	SortID    int `json:"srotId"`
}

// Labels struct
type Labels struct {
	LabelID          int    `json:"labelId"`
	LabelName        string `json:"labelName"`
	LabelAlias       string `json:"labelAlias"`
	LabelDescription string `json:"labelDescription"`
}

// ArticleLabel struct
type ArticleLabel struct {
	ArticleID int `json:"articleId"`
	LabelID   int `json:"labelId"`
}
