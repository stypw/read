package main

import (
	"database/sql"
	router "rest/router"
)

var (
	routers  map[string]*router.Router = make(map[string]*router.Router)
	globalDb *sql.DB                   = nil
)
