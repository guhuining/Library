package server

import (
	"library/data"
	"library/tools"
	"net/http"
)

// @title		LoginLibrarian
// @description	图书管理员登陆
// @param		w	http.ResponseWriter
// @param		r	*http.Request
func LoginLibrarian(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	var librarian = &data.Librarian{UserName: postData["UserName"].(string)}
	err = librarian.RetrieveByUserName()
	// 如果登录成功，设置session
	if librarian.Password == postData["Password"].(string) {
		session, _ := store.Get(r, "library")
		session.Values["LibrarianID"] = librarian.LibrarianID
		session.Values["UserName"] = librarian.UserName
		session.Values["Roll"] = "Librarian" // 设置权限
		err = session.Save(r, w)

		if err != nil { // session写入失败，登陆失败
			w.Write(tools.ApiReturn(1, "登录失败", nil))
			return
		}
		w.Write(tools.ApiReturn(0, "登录成功", nil))
		return
	} else {
		w.Write(tools.ApiReturn(1, "密码错误", nil))
		return
	}
}
