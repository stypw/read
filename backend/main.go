package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"io/ioutil"

	JSON "rest/json"

	router "rest/router"
)

var routers map[string]*router.Router = make(map[string]*router.Router)

func main() {

	ct, er := ioutil.ReadFile("./config.json")
	if er != nil {
		fmt.Println("config file open field")
		return
	}

	obj, er := JSON.FromString(string(ct))
	if er != nil {
		fmt.Println("config fmt error")
		return
	}

	mysql, y := obj["mysql"]
	if !y {
		fmt.Println("config error,have not mysql field!")
		return
	}
	mysqlString, _ := JSON.AsString(mysql)
	if mysqlString == "" {
		fmt.Println("config field mysql is empty")
		return
	}

	listen, y := obj["listen"]
	if !y {
		fmt.Println("config error,have not listen field!")
		return
	}
	listenString, _ := JSON.AsString(listen)
	if listenString == "" {
		fmt.Println("config field listen is empty")
		return
	}

	db, err := sql.Open("mysql", mysqlString)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	for _, rt := range routers {
		rt.Start(db)
	}

	http.ListenAndServe(listenString, nil)

	fmt.Println("Read Server Exited!")
}
