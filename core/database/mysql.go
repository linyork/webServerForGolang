package database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/xormplus/xorm"
    "gin/core/config"
    "log"
)

var mysql *xorm.Engine

func init() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci",
        config.DB.User,
        config.DB.Password,
        config.DB.Host,
        config.DB.Port,
        config.DB.DataBase)

    log.Println("DB dsn:", dsn)

    var err error
    mysql, err = xorm.NewEngine("mysql", dsn)
    if err != nil {
        log.Fatalln(err)
    }

    mysql.SetMaxIdleConns(1)
    mysql.SetMaxOpenConns(3)

    mysql.ShowSQL(true)

    Ping()
}

func Mysql() *xorm.Engine {
    return mysql
}

func Ping() {
    if err := mysql.Ping(); err != nil {
        log.Println("Connect to mysql failed.")
    } else {
        log.Println("Connect to mysql successfully.")
    }
}
