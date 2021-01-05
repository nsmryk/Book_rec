package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "books/service"
    "books/model"
    "strconv"
)


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

func BookUpdate(c *gin.Context){
    book := model.Book{}
    err := c.Bind(&book)
    if err != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    bookService :=service.BookService{}
    err = bookService.UpdateBook(&book)
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
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