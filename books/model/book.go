package model

type Book struct {
	Id    int64  `xorm:"pk" form:"id" json:"id"`
	Title string `form:"title" json:"title"`
	Score int64  `form:"score" json:"score"`
	Memo string `form:"memo" json:"memo"`
	
}