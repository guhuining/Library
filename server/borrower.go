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

func LoginBorrower(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	borrower := &data.Borrower{
		UserName: postData["UserName"].(string),
	}
	err = borrower.RetrieveBorrowerByUserName()

	if err != nil { // 数据库错误
		w.Write(tools.ApiReturn(1, err.Error(), nil))
	} else { //获取到数据
		if borrower.Password != postData["Password"] { //密码错误
			w.Write(tools.ApiReturn(1, "账号或密码错误", nil))
			return
		} else { // 登陆成功， 设置session
			session, err := store.Get(r, "library")
			if err != nil {
				w.Write(tools.ApiReturn(1, "服务器错误", nil))
				return
			}
			session.Values["UID"] = borrower.UID
			if borrower.Card.CardNO == "" {
				session.Values["CardNO"] = nil
			} else {
				session.Values["CardNO"] = borrower.Card.CardNO
				session.Values["Name"] = borrower.Card.Name
				session.Values["Major"] = borrower.Card.Major
				session.Values["BorrowerType"] = borrower.Card.BorrowerType.BorrowerType
			}
			err = session.Save(r, w)
			if err != nil {
				w.Write(tools.ApiReturn(1, "登陆失败", nil))
				return
			}
			w.Write(tools.ApiReturn(0, "登录成功", nil))
		}
	}
}

func GetPublicationByName(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	session, err := store.Get(r, "library")
	if err != nil {
		return
	}
	if _, ok := session.Values["UID"]; !ok {
		w.Write(tools.ApiReturn(1, "请先登录", nil))
		return
	}
	publication := &data.Publication{
		Name: postData["Name"].(string),
	}
	publications, err := publication.RetrieveByName()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "查询成功", &map[string]interface{}{"Publications": publications}))
	}
}

// @title	OrderPublication
// @description	预定图书
func OrderPublication(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	session, err := store.Get(r, "library")
	if err != nil {
		return
	}
	if _, ok := session.Values["UID"]; !ok {
		w.Write(tools.ApiReturn(1, "请先登录", nil))
		return
	}
	orderItem := &data.OrderItem{
		Publication: data.Publication{
			PublicationID: int64(postData["PublicationID"].(float64)),
		},
		Card: data.Card{
			CardNO: session.Values["CardNO"].(string),
		},
	}
	err = orderItem.Insert()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "预定成功", nil))
	}
}
