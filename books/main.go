package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "books/controller"
    "books/service"
	"books/middleware"
    "net/http"
    "strconv"
    "time"
    //"fmt"
)

func main() {
	engine:= gin.Default()
    engine.Static("/static", "./static")
    engine.LoadHTMLGlob("templates/*.html")
    // ミドルウェア
    engine.Use(middleware.RecordUaAndTime)
    
    engine.GET("/", func(c *gin.Context) {
        
        ctrl := service.BookService{}
        result := ctrl.GetBookList()
        total, month := controller.BookCount(c)
        books := controller.BookRank(c)
        c.HTML(http.StatusOK, "index.html", gin.H{
             // htmlに渡す変数を定義
            "result": result,
            "month": month,
            "total": total,
            "ranking":books,
        })
    })
    engine.GET("/book/get/:id", func(c *gin.Context) {
        id := c.Param("id")
        intId, _ := strconv.ParseInt(id, 10, 64)
        result := controller.BookGet(intId)
        date := result.Date
        year := date.Year()
        month := date.Month()
        c.HTML(http.StatusOK, "detail.html", gin.H{
            "result": result,
            "year":year,
            "month":month,
        })
    })
    // CRUD 
    engine.GET("/book/new", func(c *gin.Context) {
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "add.html", gin.H{
            "message":"Add New Book",
        })
    })
    engine.GET("/book/update/:id", func(c *gin.Context) {
        id := c.Param("id")
        intId, _ := strconv.ParseInt(id, 10, 64)
        datestr := c.Param("date")
        date, _ := time.Parse(datestr,"")
        c.HTML(http.StatusOK, "update.html", gin.H{
            "id": intId,
            "date": date,
        })
    })
    bookEngine := engine.Group("/book")
    {
        v1 := bookEngine.Group("/v1")
        {
            v1.POST("/add", func(c *gin.Context) {
                controller.BookAdd(c)
                c.HTML(http.StatusOK, "add.html", gin.H{
                    "message":"Back to home page",
                })
            })
            v1.GET("/list", controller.BookList)
            v1.POST("/update/:id", func(c *gin.Context) {
                id := c.Param("id")
                intId, _ := strconv.ParseInt(id, 10, 64)
                title := c.PostForm("title")
                scorestr := c.PostForm("score")
                score, _ := strconv.ParseInt(scorestr, 10, 64)
                memo := c.PostForm("memo")
                datestr := c.PostForm("date")
                date, _ := time.Parse(datestr,"")
                c.HTML(http.StatusOK, "update.html", gin.H{
                    "message": "Changed your Data.",
                    "id" : intId,
                })
                controller.BookUpdate(intId,title,score,memo,date,c)
                
            })
            v1.POST("/delete", func(c *gin.Context) {
                controller.BookDelete(c)
                c.HTML(http.StatusOK, "delete.html", gin.H{})
            })
        }
    }
    engine.Run(":8080")
}