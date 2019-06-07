package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db sql.DB

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
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER PRIMARY KEY , url varchar , shorturl varchar , redircount INTEGER )")
	statement.Exec()
	statement, err := database.Prepare("INSERT INTO URLs (url, shorturl, redircount) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	//statement.Exec("gogo.com", "go.com", 0)
	//st, _ := database.Query("select shorturl from urls")
	//var shr string
	//for st.Next() {
	//	st.Scan(&shr)
	//	fmt.Println(shr)
	//}
	return database

}

//second func : insert to data base if is not exist
func insertdb(id int, inurl string) string {
	var res string
	stQuery := "SELECT shorturl from urls where url =" + "'" + inurl + "'"
	statement, err := db.Query(stQuery)
	if err != nil {
		fmt.Println(err)
	} else {
		for statement.Next() {
			scanerr := statement.Scan(&res)
			if scanerr != nil {
				fmt.Println(scanerr)
			}
			fmt.Println(res)
		}
	}
	if err != nil {
		fmt.Println("has no same url")
		res = createshorturl(id)
		st, error := db.Prepare("insert into urls (id,url, shorturl, redircount) values (?,?, ?, ?)")
		st.Exec(id, inurl, res, 0)
		if error != nil {
			fmt.Println(error)
		}

	}

	return res
}

func main() {
	inurl := getURL()
	db = *createdb()
	shorturl := insertdb(10, inurl)
	fmt.Println(shorturl)
	createdb()

}
