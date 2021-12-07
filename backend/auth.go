package main

import (
	"errors"
	"net/http"
	JSON "rest/json"
)

const (
	cookieKey = "read.glng.xyz-sign"
)

func auth(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error) {
	ck, err := req.Cookie(cookieKey)
	if err != nil {
		return 1, nil, err
	}
	if ck == nil {
		ck = &http.Cookie{Name: cookieKey, Value: "hello world"}
	}

	return 1, nil, errors.New("not cookie")
}
