package main

import (
	router "rest/router"
)

var cityRouter = router.Router{
	Pattern: "/api/city",
	Table:   "tb_city",
	Before:  auth,
}

func init() {
	routers["/api/city"] = &cityRouter
}
