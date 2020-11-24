package main

import (
	"16_DB/DB"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "Bank")
	if err != nil{
		log.Fatal("DB is not connected", err)
	}

	err = DB.DbInit(db)
	err = DB.DbInitCurrency(db)
	err = DB.DbInitAccaunts(db)

	if err != nil {
		log.Fatal("error ", err)
	}
	fmt.Println("ok")
}

//все 3 таблицы реализовать
// select, insert, update, delete, drop - read