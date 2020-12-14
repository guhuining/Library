package server

import (
	"library/data"
	"library/tools"
	"net/http"
	"strings"
)

// @title	AddPublicationType
// @description	添加出版物类型
// @param	w	http.ResponseWriter
// @param	r	*http.Request
func AddPublicationType(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	// 鉴权
	session, err := store.Get(r, "library")
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	}
	if session.Values["Roll"] != "Administrator" {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	publicationType := &data.PublicationType{
		PublicationType: postData["PublicationType"].(string),
		Fine:            int64(postData["Fine"].(float64)),
	}
	err = publicationType.Insert()
	if err != nil && strings.Index(err.Error(), "Duplicate") != -1 {
		w.Write(tools.ApiReturn(1, "已有该出版物类型", nil))
	} else if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "添加成功", nil))
	}
}
