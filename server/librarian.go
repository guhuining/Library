package server

import (
	"library/data"
	"library/my_error"
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

// @title	IsOutOfTime
// @description	查询是否逾期
func IsOutOfTime(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	if err != nil {
		return
	}
	// 鉴权
	ok, err := authorizeLibrarian(r)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else if !ok {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	borrowItem := &data.BorrowItem{
		BorrowItemID: int64(postData["BorrowItemID"].(float64)),
	}
	err = borrowItem.IsOutOfTime()
	if err != nil && err.Error() == my_error.BorrowOutOfTimeError.Error() { // 有图书逾期未还
		err = borrowItem.GetFine()
		if err != nil {
			w.Write(tools.ApiReturn(1, "服务器错误", nil))
			return
		}

		ret := &map[string]interface{}{
			"Fine": borrowItem.Publication.PublicationType.Fine,
		}
		w.Write(tools.ApiReturn(-1, "已逾期，需缴纳罚金", ret))
	} else if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else {
		w.Write(tools.ApiReturn(0, "未逾期", nil))
	}
}
