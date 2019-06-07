package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
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
	return database

}

//second func : insert to data base if is not exist
func insertdb(id int, inurl string) string {
	var res string
	sqlQuery := "SELECT shorturl FROM urls WHERE url =" + inurl + ""
	err := db.QueryRow(sqlQuery).Scan(&res)
	fmt.Println(res)
	//}
	if err != nil {
		fmt.Println("has no same url")
		res = createshorturl(id)
		sqlQuery = "INSERT INTO urls (id, url, shorturl, redircount) values (" + strconv.Itoa(id) + "," + inurl + "," + res + "0)"
		//st, error := db.Prepare("INSERT INTO urls (url, shorturl, redirvount) values (?, ?, ?)")
		//st.Exec(inurl, res, 0)
		st, error := db.Prepare(sqlQuery)
		if error != nil {
			panic(error)
		}
		st.Exec()

		if error != nil {
			println("injaaaaaaaaaaa")
		}

	}

	return res
}

func main() {
	inurl := getURL()
	db = *createdb()
	shorturl := insertdb(1000, inurl)
	fmt.Println(shorturl)
	createdb()

}
