package server

import (
	"library/data"
	"library/tools"
	"net/http"
	"strings"
)

// @title	CreateBorrower
// @description	注册借阅者
func CreateBorrower(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	borrower := &data.Borrower{
		UserName: postData["UserName"].(string),
		Password: postData["Password"].(string),
	}
	err = borrower.Create()
	if err != nil && strings.Index(err.Error(), "Duplicate entry") != -1 {
		w.Write(tools.ApiReturn(1, "账号已被注册", nil))
	} else if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "注册成功", nil))
	}
}
