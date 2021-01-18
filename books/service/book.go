package service

import (
    "books/model"
    "time"
    "fmt"
)

type BookService struct {}
/*
func  NewBook() BookService {
    return BookService{}
}
*/
func (BookService) SetBook(book *model.Book) error {
    _, err := DbEngine.Insert(book)
    if err!= nil{
         return  err
    }
    return nil
}

func (BookService) GetById(id int64) model.Book {
    res := model.Book{}
    _ ,err := DbEngine.ID(id).Get(&res) 
    if err != nil {
        fmt.Printf("error: %s",err)
    } 
    return res
}

func (BookService) GetBookList() []model.Book {
    tests := make([]model.Book, 0)
    err := DbEngine.Distinct("id", "title", "score", "memo","date").Limit(10, 0).Find(&tests)
    if err != nil {
        panic(err)
    }
    return tests
}
func (BookService) CountBooks(date time.Time) int64 {   
    book := new(model.Book)
    t_bgn := time.Date(date.Year(),date.Month(),1, 0, 0, 0,    000000000, time.Local)
    t_end := time.Date(date.Year(),date.Month(),1, 23, 59, 59, 999999999, time.Local).AddDate(0, 1, -1)
    total, _ := DbEngine.Where("date >=? AND date <=?",t_bgn ,t_end).Count(book)
    return total
}
func (BookService) GetBookRanking() []model.Book {
    book := make([]model.Book, 0)
    DbEngine.Desc("score").Asc("title").Limit(5, 0).Find(&book)
    return book
}
func (BookService) UpdateBook(id int64,title string, score int64, memo string,date time.Time) error {
    newBook := new(model.Book)
    newBook.Id = id
    newBook.Title = title
    newBook.Score = score
    newBook.Memo = memo
    newBook.Date = date
    _, err := DbEngine.Id(id).Update(newBook)
    if err != nil {
        return err
    }
    return nil
}

func (BookService) DeleteBook(id int) error {
    book := new(model.Book)
    _, err := DbEngine.Id(id).Delete(book)
    if err != nil{
        return err
    }
    return nil
}