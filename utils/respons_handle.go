/*
* @Author: zh
* @Email: zhangheng9394@163.com
* @Date: 2020/12/3 16:38
 */
package utils

import (
	"gin_web/models"
	"net/http"
)

func ResponseHttpCode(err interface{}) (httCode int) {

	code := new(int)
	if err != nil {
		*code = http.StatusInternalServerError
	} else {
		*code = http.StatusOK
	}
	return *code
}

func ResponseData(err interface{}) models.ResData {
	/*
	对返回值进行封装
		{
			"Status": 1,
			"Msg": "success",
			"Result": null
		}
	 */
	code := new(int)
	msg := new(string)
	if err != nil {
		*code = 0
		*msg = "failure"
	} else {
		*code = 1
		*msg = "success"
	}
	var resp models.ResData
	resp.Status = *code
	resp.Msg = *msg

	return resp
}
