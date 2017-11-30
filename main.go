package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io/ioutil" //baca file stopwordbahasa.csv
	"log"
	"reflect" //fungsi in_array
	"regexp"
	"strings"
	"time"

	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//getScrapeKompas()
	//getScrapeDetik()
	build()
}

func build() {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_PROTOCOL+"("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset="+DB_CHARSET)
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}
	defer db.Close()

	///stopword
	content, err := ioutil.ReadFile("stopwordbahasa.csv")
	check(err)
	stopword := strings.Split(string(content), "\n")
	///
	fixer, err := regexp.Compile("[^a-zA-Z0-9]+")
	check(err)
	var title string
	rows, err := db.Query("select title from news")
	check(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&title)
		check(err)
		//log.Println(title)
		a := strings.Split(title, " ")
		//fl := 1
		for k, val := range a {
			fmt.Println(val)
			val := strings.ToLower(fixer.ReplaceAllString(val, ""))
			if in_array(val, stopword) {
				unset(a, k)
			}
		}
	}
	err = rows.Err()
	check(err)
}

func unset(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func in_array(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}

///////
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// remove whitespaces
func rws(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// generate sha1
func sha1gen(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// DB
const (
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_NAME     = "news"
	DB_USER     = "ammar"
	DB_PASS     = "858869123"
	DB_CHARSET  = "utf8"
	DB_PROTOCOL = "tcp"
)

func insertDB(hash, site, title string, created_at time.Time) {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_PROTOCOL+"("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset="+DB_CHARSET)
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}

	stmt, err := db.Prepare("INSERT IGNORE INTO news(hash, site, title, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	//res, err := stmt.Exec(hash, site, title, created_at)
	_, err = stmt.Exec(hash, site, title, created_at)
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	defer db.Close()

	fmt.Printf("Tersimpan - %s\n", title)
	/*
		id, _ := res.LastInsertId()
		row, _ := res.RowsAffected()
		fmt.Printf("Inserted row: %d - %d\n", id, row)
	*/
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
