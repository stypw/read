package main

import (
	router "rest/router"
)

var cityRouter = router.Router{
	Pattern: "/api/city",
	Table:   "tb_city",
	Before:  isAuth,
}

func init() {
	routers["/api/city"] = &cityRouter
}
