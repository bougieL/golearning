package utils

import "net/http"

// Res detail
func Res(code int, data interface{}, message ...string) (int, map[string]interface{}) {
	var res map[string]interface{}
	res = make(map[string]interface{})
	res["code"] = code
	if len(message) > 0 {
		res["message"] = message[0]
	} else {
		res["message"] = http.StatusText(code)
	}
	res["data"] = data
	return code, res
}
