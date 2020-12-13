package tools

import "encoding/json"

// @title apiReturn
// @description 将数据编码为json格式，用来返回给前端
// @param	code		int						"状态码，-1：服务器错误；0：正确；1：数据传输错误"
// @param	msg			string					"说明信息"
// @param	data		map[string]interface{}	"需要传输的数据"
// @return	jsonString	[]byte					"编码后的json字节slice"
func ApiReturn(code int, msg string, data *map[string]interface{}) []byte {
	var jsonMap = map[string]interface{}{}
	if data == nil {
		jsonMap["code"] = code
		jsonMap["msg"] = msg
		jsonMap["data"] = nil
	} else {
		jsonMap["code"] = code
		jsonMap["msg"] = msg
		jsonMap["data"] = *data
	}

	jsonString, err := json.MarshalIndent(jsonMap, "", "    ")
	if err != nil {
		return []byte(`{"code": -1, "msg":"server error", “data”: ""}`)
	}
	return jsonString
}
