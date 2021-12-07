package main

import "net/http"

type listener struct {
	patterns []string
	handle   http.HandlerFunc
}

var listeners = make(map[string]*listener)
