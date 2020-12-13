package server

import (
	"library/data"
	"library/tools"
	"net/http"
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
	// TODO 写好登录后添加鉴权
	var administrator = &data.Administrator{UserName: postData["UserName"].(string), Password: postData["Password"].(string)}
	err = administrator.Create() // 添加系统管理员
	if err != nil {
		w.Write(tools.ApiReturn(1, "添加失败", nil))
		print(err.Error())
		return
	}
	w.Write(tools.ApiReturn(0, "添加成功", nil))
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
		session.Values["Roll"] = "administrator" // 设置权限
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
