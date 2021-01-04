package service

import (
    "books/model"
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


func (BookService) GetBookList() []model.Book {
    tests := make([]model.Book, 0)
    err := DbEngine.Distinct("id", "title", "score", "memo").Limit(10, 0).Find(&tests)
    if err != nil {
        panic(err)
    }
    return tests
}

func (BookService) UpdateBook(newBook *model.Book) error {
    _, err := DbEngine.Id(newBook.Id).Update(newBook)
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