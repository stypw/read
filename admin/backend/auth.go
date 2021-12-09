package main

import (
	"errors"
	"net/http"
	JSON "rest/json"
	"rest/orm"
)

const (
	cookieKey = "read.glng.xyz-sign"
)

type user struct {
	acc      string
	nickname string
}

var users = make(map[string]*user, 0)

func isAuth(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error) {
	ck, err := req.Cookie(cookieKey)
	if err != nil {
		return 1, nil, errors.New("cookie错误，请设置浏览器应许cookie")
	}
	if ck == nil {
		return 2, nil, errors.New("请先登录")
	}

	user := users[ck.Value]
	if user == nil {
		return 2, nil, errors.New("请先登录")
	}

	return 0, nil, nil
}

func doLogin(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error) {
	body, err := JSON.FromStream(req.Body)

	if err != nil || body == nil {
		return 1, nil, errors.New("用户名密码不能为空")
	}

	var acc = JSON.GetString(JSON.GetProperty(body, "acc"))
	var pwd = JSON.GetString(JSON.GetProperty(body, "pwd"))
	if acc == "" || pwd == "" {
		return 1, nil, errors.New("用户名密码不能为空")
	}
	var pwdSha512 = toSha512String(pwd)

	orm := orm.Orm{TableName: "tb_admin", Db: globalDb}

	item, err := orm.First(JSON.Object{
		"acc": JSON.String(acc),
		"pwd": JSON.String(pwdSha512),
	}, nil)
	if err != nil {
		return 1, nil, errors.New("数据库执行错误")
	}
	if item == nil {
		return 1, nil, errors.New("用户名或密码错误")
	}

	accSha512 := toSha512String(acc)
	users[accSha512] = &user{acc: acc, nickname: JSON.GetString(JSON.GetProperty(item, "nickname"))}
	http.SetCookie(w, &http.Cookie{Name: cookieKey, Value: accSha512})
	return 0, JSON.Object{
		"acc":      JSON.String(acc),
		"nickname": JSON.GetProperty(item, "nickname"),
	}, nil
}

func login(w http.ResponseWriter, req *http.Request) {
	code, ret, err := doLogin(w, req)
	var responseData JSON.Object = make(JSON.Object)

	responseData["code"] = JSON.Number(float64(code))
	if err != nil {
		responseData["error"] = JSON.String(err.Error())
	} else {
		responseData["data"] = ret
	}
	w.Write([]byte(responseData.ToString()))
}

func doAuth(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error) {
	ck, err := req.Cookie(cookieKey)
	if err != nil {
		return 1, nil, errors.New("cookie错误，请设置浏览器应许cookie")
	}
	if ck == nil {
		return 2, nil, errors.New("未登录")
	}

	user := users[ck.Value]
	if user == nil {
		return 2, nil, errors.New("未登录")
	}

	return 0, JSON.Object{
		"acc":      JSON.String(user.acc),
		"nickname": JSON.String(user.nickname),
	}, nil
}

func auth(w http.ResponseWriter, req *http.Request) {
	code, ret, err := doAuth(w, req)
	var responseData JSON.Object = make(JSON.Object)

	responseData["code"] = JSON.Number(float64(code))
	if err != nil {
		responseData["error"] = JSON.String(err.Error())
	} else {
		responseData["data"] = ret
	}
	w.Write([]byte(responseData.ToString()))
}

func init() {
	listeners["login"] = &listener{patterns: []string{"/api/login", "/api/login/"}, handle: login}
	listeners["auth"] = &listener{patterns: []string{"/api/auth", "/api/auth/"}, handle: auth}
}
