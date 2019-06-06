package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
//now this is not good func
//daghoonesh kardam ke ye insert dorost dashte basham
func createdb() *sql.DB {
	database, _ := sql.Open("sqlite3", "./data.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER PRIMARY KEY , url TEXT, shorturl TEXT, redircount INTEGER )")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO URLs (url, shorturl, redirvount) VALUES (?, ?, ?)")
	statement.Exec("gogo.com", "go.com", 0)
	st, _ := database.Query("select shorturl from urls")
	var shr string
	for st.Next() {
		st.Scan(&shr)
		fmt.Println(shr)
	}
	return database

}

//second func : insert to data base if is not exist
func insertdb(id int, inurl string, db *sql.DB) string {
	var res string
	var id2 int

	statement, err := db.Query("SELECT id from URLs")

	statement.Scan(&id2)
	fmt.Println(id2)
	//}
	if err == sql.ErrNoRows {
		fmt.Println("has no same url")
		res = createshorturl(id)
		st, error := db.Prepare("insert into URLs (url, shorturl, redirvount) values (?, ?, ?)")
		st.Exec(inurl, res, 0)
		if error != nil {
			println("injaaaaaaaaaaa")
		}
		//	return res
	}

	//println(res)
	return strconv.Itoa(id2)
}

func main() {
	//inurl := getURL()
	//db := createdb()
	//shorturl := insertdb(1000, inurl, db)
	//fmt.Println(shorturl)
	createdb()

}
