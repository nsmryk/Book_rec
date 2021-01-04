package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "books/controller"
    "books/service"
	"books/middleware"
	"net/http"
)
func main() {
	engine:= gin.Default()

    engine.LoadHTMLGlob("templates/*.html")
    engine.GET("/", func(c *gin.Context) {
        ctrl := service.BookService{}
        result := ctrl.GetBookList()
        c.HTML(http.StatusOK, "index.html", gin.H{
             // htmlに渡す変数を定義
            "response": result,
        })
    })
    // ミドルウェア
    engine.Use(middleware.RecordUaAndTime)
    // CRUD 書籍
    bookEngine := engine.Group("/book")
    {
        v1 := bookEngine.Group("/v1")
        {
            v1.POST("/add", controller.BookAdd)
            v1.GET("/list", controller.BookList)
            v1.PUT("/update", controller.BookUpdate)
            v1.DELETE("/delete", controller.BookDelete)
        }
    }
    engine.Run(":8080")
}