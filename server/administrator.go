package server

import (
	"library/data"
	"library/tools"
	"net/http"
	"strings"
)

// @title		CreateAdministrator
// @description	添加系统管理员
// @param		w	http.ResponseWriter
// @param		r	*http.Request
func CreateAdministrator(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r) // 获取请求数据
	if err != nil {
		return
	}
	session, _ := store.Get(r, "library")
	if session.Values["Roll"] != "Administrator" {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	var administrator = &data.Administrator{UserName: postData["UserName"].(string), Password: postData["Password"].(string)}
	err = administrator.Create() // 添加系统管理员
	if err != nil && strings.Index(err.Error(), "duplicate") != -1 {
		w.Write(tools.ApiReturn(1, "该账号已存在", nil))
	} else if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "添加成功", nil))
	}
}

// @title		LoginAdministrator
// @description	系统管理员登陆
// @param		w	http.ResponseWriter
// @param		r	*http.Request
func LoginAdministrator(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	var administrator = &data.Administrator{UserName: postData["UserName"].(string)}
	err = administrator.RetrieveByUserName()
	// 如果登录成功，设置session
	if administrator.Password == postData["Password"].(string) {
		session, _ := store.Get(r, "library")
		session.Values["AdministratorID"] = administrator.AdministratorID
		session.Values["UserName"] = administrator.UserName
		session.Values["Roll"] = "Administrator" // 设置权限
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

// @title	CreateLibrarian
// @description	创建图书管理员
// @param		w	http.ResponseWriter
// @param		r	*http.Request
func CreateLibrarian(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r) // 获取请求数据
	if err != nil {
		return
	}
	//鉴权
	ok, err := authorizeAdministrator(r)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else if !ok {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	var librarian = &data.Librarian{UserName: postData["UserName"].(string), Password: postData["Password"].(string)}
	err = librarian.Create() // 添加系统管理员
	if err != nil && strings.Index(err.Error(), "duplicate") != -1 {
		w.Write(tools.ApiReturn(1, "该账号已存在", nil))
	} else if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "添加成功", nil))
	}

}

// @title	DeleteLibrarian
// @description	删除图书管理员
// @param		w	http.ResponseWriter
// @param		r	*http.Request
func DeleteLibrarian(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r) // 获取请求数据
	if err != nil {
		return
	}
	//鉴权
	ok, err := authorizeAdministrator(r)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else if !ok {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	librarian := &data.Librarian{
		LibrarianID: int64(postData["LibrarianID"].(float64)),
	}
	err = librarian.Delete()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "删除成功", nil))
	}
}
