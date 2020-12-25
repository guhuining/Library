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

// @title	BorrowPublication
// @description	借阅图书
func BorrowPublication(w http.ResponseWriter, r *http.Request) {
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
		Card: data.Card{
			CardNO: postData["CardNO"].(string),
		},
		Publication: data.Publication{
			PublicationID: int64(postData["PublicationID"].(float64)),
		},
	}
	err = borrowItem.Borrow()
	if err != nil {
		if err.Error() == my_error.BorrowOutOfTimeError.Error() {
			w.Write(tools.ApiReturn(1, "还有图书逾期未归还", nil))
		} else if err.Error() == my_error.InventoryNotEnoughError.Error() {
			w.Write(tools.ApiReturn(1, "库存不足", nil))
		} else if err.Error() == my_error.MaxBorrowNumberError.Error() {
			w.Write(tools.ApiReturn(1, "借阅量已达上限", nil))
		} else {
			w.Write(tools.ApiReturn(1, err.Error(), nil))
		}
	} else {
		w.Write(tools.ApiReturn(0, "借阅成功", nil))
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

// @title	ReturnPublication
// @description	还书
func ReturnPublication(w http.ResponseWriter, r *http.Request) {
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
	err = borrowItem.Return()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "还书成功", nil))
	}
}

// @title	ReturnPublication
// @description	获取所有借阅订单
func GetBorrowItem(w http.ResponseWriter, r *http.Request) {
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
		Card: data.Card{
			CardNO: postData["CardNO"].(string),
		},
	}
	borrowItems, err := borrowItem.GetBorrowItem()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	}
	var ret []map[string]interface{}
	for _, item := range borrowItems {
		ret = append(ret, map[string]interface{}{
			"BorrowItemID": item.BorrowItemID,
			"Name":         item.Publication.Name,
			"Author":       item.Publication.Author,
		})
	}
	w.Write(tools.ApiReturn(0, "获取数据成功", &map[string]interface{}{"BorrowItems": ret}))
}

// @title LibrarianGetOrderItem
// @description 获取所有未兑现订单
func LibrarianGetOrderItem(w http.ResponseWriter, r *http.Request) {
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
	orderItem := &data.OrderItem{
		Card: data.Card{
			CardNO: postData["CardNO"].(string),
		},
	}
	results, err := orderItem.RetrieveByCardNO()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		var ret []map[string]interface{}
		for _, result := range results {
			temp := &map[string]interface{}{
				"OrderItemID":   result.OrderItemID,
				"PublicationID": result.Publication.PublicationID,
				"Name":          result.Publication.Name,
				"Author":        result.Publication.Author,
				"Total":         result.Publication.Total,
				"Inventory":     result.Publication.Inventory,
			}
			ret = append(ret, *temp)
		}
		w.Write(tools.ApiReturn(0, "获取成功", &map[string]interface{}{"OrderItem": ret}))
	}
}

func LibrarianGetPublicationByName(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	// 鉴权
	ok, err := authorizeLibrarian(r)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else if !ok {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
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

func OrderBorrow(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
	// 鉴权
	ok, err := authorizeLibrarian(r)
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
		return
	} else if !ok {
		w.Write(tools.ApiReturn(1, "权限不足", nil))
		return
	}
	orderItem := &data.OrderItem{
		OrderItemID: int64(postData["OrderItemID"].(float64)),
	}
	err = orderItem.Borrow()
	if err != nil && err.Error() == my_error.BorrowOutOfTimeError.Error() {
		w.Write(tools.ApiReturn(1, "有超期图书未归还", nil))
	} else if err != nil && err.Error() == my_error.MaxBorrowNumberError.Error() {
		w.Write(tools.ApiReturn(1, "超出借阅上限", nil))
	} else if err != nil && err.Error() == my_error.InventoryNotEnoughError.Error() {
		w.Write(tools.ApiReturn(1, "库存不足", nil))
	} else if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "借阅成功", nil))
	}
}

func GetPrice(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
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
	err = borrowItem.GetPrice()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "获取成功", &map[string]interface{}{"Price": borrowItem.Publication.Price}))
	}
}

func Lost(w http.ResponseWriter, r *http.Request) {
	postData, err := tools.GetPostBody(w, r)
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
	err = borrowItem.Lost()
	if err != nil {
		w.Write(tools.ApiReturn(1, "服务器错误", nil))
	} else {
		w.Write(tools.ApiReturn(0, "确认丢失", nil))
	}
}
