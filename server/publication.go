package server

import (
	"library/data"
	"library/tools"
	"net/http"
)

// @title	AddPublication
// @description	添加出版物
// @param	w	http.ResponseWriter
// @param	r	*http.Request
func AddPublication(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	// 鉴权
	ok, err := authorizeAdministrator(r)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else if !ok {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	publication := &data.Publication{
		Name:            postData["Name"].(string),
		ISBN:            postData["ISBN"].(string),
		Price:           int64(postData["Price"].(float64)),
		Total:           int64(postData["Total"].(float64)),
		Inventory:       int64(postData["Total"].(float64)),
		PublicationType: data.PublicationType{PublicationType: postData["PublicationType"].(string)},
		Author:          postData["Author"].(string),
	}
	err = publication.Insert()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "添加成功", nil))
	}
}

// @title	DeletePublication
// @description	删除出版物
// @param	w	http.ResponseWriter
// @param	r	*http.Request
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	// 鉴权
	ok, err := authorizeAdministrator(r)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else if !ok {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	publication := &data.Publication{
		PublicationID: int64(postData["PublicationID"].(float64)),
	}
	err = publication.Delete()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "删除成功", nil))
	}
}
