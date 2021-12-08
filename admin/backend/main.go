package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"io/ioutil"

	JSON "rest/json"
)

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

	mysql := JSON.GetString(JSON.GetProperty(obj, "mysql"))
	if mysql == "" {
		fmt.Println("config error,have not mysql field!")
		return
	}

	listen := JSON.GetString(JSON.GetProperty(obj, "listen"))
	if listen == "" {
		fmt.Println("config error,have not listen field!")
		return
	}

	db, err := sql.Open("mysql", mysql)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	globalDb = db

	defer db.Close()

	for _, rt := range routers {
		rt.Start(db)
	}

	listeners["/"] = &listener{patterns: []string{"/"}, handle: func(w http.ResponseWriter, req *http.Request) {
		var responseData = JSON.Object{
			"code":    JSON.Number(4),
			"message": JSON.String("not found:" + req.URL.Path + req.URL.RawQuery),
		}
		w.Write([]byte(responseData.ToString()))
	}}

	for _, ls := range listeners {
		for _, p := range ls.patterns {
			http.HandleFunc(p, ls.handle)
		}
	}

	http.ListenAndServe(listen, nil)

	fmt.Println("Read Server Exited!")
}
