package model

import (
    "time"
)
type Book struct {
	Id    int64  `xorm:"pk autoincr" form:"id" json:"id"`
	Title string `form:"title" json:"title"`
	Score int64  `form:"score" json:"score"`
	Memo string `form:"memo" json:"memo"`
	Date time.Time `xorm:"DateTime" form:"date" json:"date"`
}