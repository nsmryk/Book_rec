package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "books/service"
    "books/model"
    "strconv"
    "fmt"
)

func BookGet(c *gin.Context) *model.Book {
    id := c.PostForm("id")
    intId, _ := strconv.ParseInt(id, 10, 0)
    
    bookService :=service.BookService{}
    res := bookService.GetById(int64(intId))
    return res
}

func BookAdd(c *gin.Context) {
     book := model.Book{}
     err := c.Bind(&book)
     if err != nil{
         c.String(http.StatusBadRequest, "Bad request")
         return
     }
    bookService :=service.BookService{}
    err = bookService.SetBook(&book)
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    
}

func BookList(c *gin.Context)  {
    bookService :=service.BookService{}
	BookLists := bookService.GetBookList()
	
    c.JSONP(http.StatusOK, gin.H{
        "message": "ok",
        "data": BookLists,
	})
	
}

func BookUpdate(id int64, title string, score int64, memo string, c *gin.Context){
    bookService :=service.BookService{}
    err := bookService.UpdateBook(id,title,score,memo)
    if err != nil{
        c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
        //c.String(http.StatusInternalServerError, "Server Error")
        return 
    }
    
}

func BookDelete(c *gin.Context){
    id := c.PostForm("id")
    intId, err := strconv.ParseInt(id, 10, 0)
    if err != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    bookService :=service.BookService{}
    err = bookService.DeleteBook(int(intId))
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    
}