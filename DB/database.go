package DB

import (
	"database/sql"
	"fmt"
)

func init()  {
	fmt.Println("hello, I am test")
}

func DbInit(db *sql.DB) (err error) {
	const usersDB = `CREATE TABLE if not exists user (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text not null,
    surname TEXT NOT NULL,
    age INTEGER NOT NULL,
    sex TEXT,
    remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(usersDB)

	if err != nil {
		return err
	}
	return nil
}
func DbInitCurrency(db *sql.DB) (err error) {
	const currency = `CREATE TABLE if not exists currency (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text
)`
	_, err = db.Exec(currency)

	if err != nil {
		return err
	}
	return nil
	}

	func DbInitAccaunts(db *sql.DB) (err error){
		const accaunts  = `CREATE TABLE if not exists accaunts (
    	id integer PRIMARY KEY AUTOINCREMENT,
    	userId integer references user(id) not null ,
    	number integer,
    	amount integer,
    	currency integer references currency(id),
    	remove BOOLEAN NOT NULL DEFAULT FALSE
)`
		_, err = db.Exec(accaunts)

		if err != nil {
			return err
		}
		return nil
	}


