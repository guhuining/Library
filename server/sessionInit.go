package server

import (
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

// @title init
// @description 初始化session和随机数种子
func init() {
	key := []byte("library")
	store = sessions.NewCookieStore(key)
}
