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
