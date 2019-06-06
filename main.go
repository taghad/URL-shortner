package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

//scan URL
func getURL() string {
	var inurl string
	fmt.Scanf("%s", &inurl)
	return inurl
}
func createdb() {
	os.Create("./data.db")
	database, _ := sql.Open("sqlite3", "./data.db")
	statement, _ := database.Prepare("create table if not exists URLs(id integer primary key, url varchar, shorturl varchar ,exptime date ,redirnum int )")
	statement.Exec()

}
func main() {

}
