package tools

import (
	"encoding/json"
	"net/http"
)

// @title getPostBody
// @description 从请求body中读取数据并转换为map
// @return	postData	map[string]interface{}	"请求中的post信息"
// @return	err 		error					"错误信息"
func GetPostBody(w http.ResponseWriter, r *http.Request) (postData map[string]interface{}, err error) {
	// 获取请求body
	bodyLen := r.ContentLength
	body := make([]byte, bodyLen)
	_, err = r.Body.Read(body)
	// 将请求body中json格式的数据读取到map中
	err = json.Unmarshal(body, &postData)
	if err != nil {
		w.Write(ApiReturn(-1, "json form error", nil))
		return
	}
	return
}
