package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var db sql.DB

//going to get shorturl & redirect to url
func redirect(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path
	println("url :", s)
	st, err := db.Query("select url from urls where shorturl = ?", "http://localhost:8161"+s)
	var res string
	if err != nil {
		fmt.Println(err)
	} else {
		for st.Next() {
			scanerr := st.Scan(&res)
			if scanerr != nil {
				fmt.Println(scanerr)
			}
		}
	}
	fmt.Println(s)
	http.Redirect(w, r, res, 301)
}

//scan URL
func getURL() string {
	var inurl string
	fmt.Scanf("%s", &inurl)
	return inurl
}

//this is for check genrated short url is exist or not
func ifexistShorturl(shorturl string) bool {
	var result string
	stQuery := "SELECT shorturl FROM urls WHERE shorturl =" + "'" + shorturl + "'"
	res, err := db.Query(stQuery)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		res.Scan(&result)
		fmt.Println(result)
	}
	//we have bug here
	if result == "" {

		//fmt.Println("voojood nadarad")
		return false
	}
	//fmt.Println("voojood darad")
	return true
}

//create shortURL
func createshorturl(id int) string {
	if !ifexistShorturl("http://localhost:8161/" + strconv.Itoa(id)) {
		return "http://localhost:8161/" + strconv.Itoa(id)
	}
	return createshorturl(rand.Int())

}

//data base func :
//first func : create data base & table
func createdb() *sql.DB {
	db, errorOpen := sql.Open("sqlite3", "./data.db")
	if errorOpen != nil {
		fmt.Println(errorOpen)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER PRIMARY KEY , url text , shorturl text , redircount INTEGER )")
	if err != nil {
		fmt.Println(err)
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println(err)
	}

	return db

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
	st, error := db.Prepare("insert into urls (url, shorturl , redircount) values (?,?,?)")

	if error != nil {
		fmt.Println(error)
	}
	_, error = st.Exec(inurl, res, 0)

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
	{
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
			http.HandleFunc("/5577006791947779410", redirect)
			err := http.ListenAndServe(":8161", nil)
			if err != nil {
				log.Fatal("ListenAndServe: ", err)
			}
			break
		case 3:
			inurl := getURL()
			db = *createdb()
			insWithCusShorturl(getURL(), inurl)
		}

	}

}
