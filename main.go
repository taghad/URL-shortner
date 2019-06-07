package main

import (
	"database/sql"
	"fmt"
	"math/rand"
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

//this is for check genrated short url is exist or not
func ifexistShorturl(shorturl string) bool {
	stQuery := "SELECT shorturl from urls where shorturl =" + "'" + shorturl + "'"
	_, err := db.Query(stQuery)
	if err != nil {
		fmt.Println(err)
		return true
	}

	return false
}

//create shortURL
func createshorturl(id int) string {
	if ifexistShorturl("taghad.gogo/" + strconv.Itoa(id)) {
		return "taghad.gogo/" + strconv.Itoa(id)
	}
	return createshorturl(rand.Int())

}

//data base func :
//first func : create data base & table
func createdb() *sql.DB {
	database, _ := sql.Open("sqlite3", "./data.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER PRIMARY KEY , url varchar , shorturl varchar,expiretime date, redircount INTEGER )")
	statement.Exec()
	statement, err := database.Prepare("INSERT INTO URLs (url, shorturl, redircount) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}

	return database

}

//second func : insert to data base & return shorturl
func insertdb(shorturl string, inurl string) string {
	var res string
	//if exist don't insert it
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
			return res
		}
	}
	//if not exist so insert

	fmt.Println(err)
	res = shorturl
	st, error := db.Prepare("insert into urls (url, shorturl,expiretime, redircount) values (?,?, DATEADD(min ,30, current_date), ?)")
	st.Exec(inurl, res, 0)

	if error != nil {
		fmt.Println(error)
	}
	return res
}

//third func : insert with custom shorturl
func insWithCusShorturl(custshort string, inurl string) bool {
	stQuery := "SELECT shorturl from urls where shorturl =" + "'" + custshort + "'"
	_, err := db.Query(stQuery)
	if err != nil {
		fmt.Println(err)
		fmt.Println("you can't do this")
		fmt.Println(insertdb(createshorturl(rand.Int()), inurl))
		return false
	} else {
		fmt.Println(insertdb(custshort, inurl))
		return true
	}

}

func main() {
	for true {
		fmt.Println("1: give shorturl \n2: redirect with shorturl \n3: set short url for your link")
		var state int
		fmt.Scanf("%d", &state)
		switch state {
		case 1:
			inurl := getURL()
			db = *createdb()
			shorturl := insertdb(createshorturl(rand.Int()), inurl)
			fmt.Println(shorturl)
			break
		case 2:
			//nothing now
			break
		case 3:
			inurl := getURL()
			db = *createdb()
			insWithCusShorturl(getURL(), inurl)
		}

	}

}
