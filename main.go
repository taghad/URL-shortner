package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
)

//scan URL
func getURL() string {
	var inurl string
	fmt.Scanf("%s", &inurl)
	return inurl
}

//create shortURL
func createshorturl(id int) string {
	return "taghad.gogo/" + strconv.Itoa(id)

}

//data base func :
//first func : create data base & table
func createdb() *sql.DB {
	os.Create("./data.db")
	database, _ := sql.Open("sqlite3", "./data.db")
	statement, _ := database.Prepare("create table if not exists URLs(id integer primary key autoincrement, url varchar , shorturl varchar , redircount int )")
	statement.Exec()
	return database

}

//second func : insert to data base if is not exist
func insertdb(id int, inurl string, db *sql.DB) string {
	var res string
	//var id1,i int
	//var urll string
	statement, err := db.Query("select shorturl from URLs where url = ?", inurl)

	for statement.Next() {
		statement.Scan(&res)
	}

	if err == sql.ErrNoRows {
		println("has no same url")
		res = createshorturl(id)
		st, error := db.Prepare("insert into URLs(url,shorturl,redirvount) values (?, ?,0)")
		st.Exec(inurl, res)
		if error != nil {
			println(error)
		}
		return res
	}

	//fmt.Println("eeee alie")
	println(res)
	return res
}

func main() {
	inurl := getURL()
	db := createdb()
	shorturl := insertdb(10, inurl, db)
	fmt.Println(shorturl)

}
