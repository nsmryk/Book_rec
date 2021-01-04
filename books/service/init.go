package service

import (
    "errors"
    "fmt"
    "github.com/go-xorm/xorm"
    "books/model"
    "log"
)

var DbEngine *xorm.Engine

func init()  {
    driverName := "mysql"
    DsName := "admin:admin@tcp(mysql)/db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    err := errors.New("")
    DbEngine, err = xorm.NewEngine(driverName,DsName)
    if err != nil && err.Error() != ""{
        log.Fatal(err.Error())
    }
    DbEngine.ShowSQL(true)
    DbEngine.SetMaxOpenConns(2)
    DbEngine.Sync2(new(model.Book))
    fmt.Println("init data base ok")
}