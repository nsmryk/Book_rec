package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "books/controller"
    "books/service"
	"books/middleware"
    "net/http"
    //"fmt"
)
func main() {
	engine:= gin.Default()

    engine.LoadHTMLGlob("templates/*.html")
    engine.GET("/", func(c *gin.Context) {
        ctrl := service.BookService{}
        result := ctrl.GetBookList()
        c.HTML(http.StatusOK, "index.html", gin.H{
             // htmlに渡す変数を定義
            "result": result,
        })
    })
    // ミドルウェア
    engine.Use(middleware.RecordUaAndTime)
    // CRUD 書籍
    engine.GET("/book/new", func(c *gin.Context) {
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "add.html", gin.H{})
    })
    bookEngine := engine.Group("/book")
    {
        v1 := bookEngine.Group("/v1")
        {
            v1.POST("/add", func(c *gin.Context) {
                controller.BookAdd(c)
                ctrl := service.BookService{}
                result := ctrl.GetBookList()
                c.HTML(http.StatusOK, "index.html", gin.H{
                    // htmlに渡す変数を定義
                    "result": result,
                })
            })
            v1.GET("/list", controller.BookList)
            v1.PUT("/update", controller.BookUpdate)
            v1.POST("/delete", func(c *gin.Context) {
                controller.BookDelete(c)
                /*ctrl := service.BookService{}
                result := ctrl.GetBookList()
                c.HTML(http.StatusOK, "index.html", gin.H{
                    // htmlに渡す変数を定義
                    "result": result,
                })*/
            })
        }
    }
    engine.Run(":8080")
}