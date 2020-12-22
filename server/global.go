package server

import (
	"library/tools"
	"net/http"
)

// @title IsLogin
// @description	判断是否登录
func IsLogin(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "library")
	if err != nil {
		w.Write(tools.ApiReturn(-1, "未登录", nil))
		return
	}
	_, ok1 := session.Values["Roll"]
	_, ok2 := session.Values["UID"]
	if ok1 || ok2 {
		w.Write(tools.ApiReturn(0, "已登录", nil))
	} else {
		w.Write(tools.ApiReturn(-1, "未登录", nil))
	}
}

// @title Logout
// @description	下线
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "library")
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	}
	for key, _ := range session.Values {
		delete(session.Values, key)
	}
	err = session.Save(r, w)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "退出成功", nil))
	}
}
