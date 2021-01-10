/*
* @Author: zh
* @Email: zhangheng9394@163.com
* @Date: 2020/12/11 14:32
 */
package test

type UserInfo struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type User struct {
	Username string //`json:"username" binding:"required"`  // 必传校验
	Password string //`json:"Password" binding:"required"`
	Data     []UserInfo `json:"data"`
}


