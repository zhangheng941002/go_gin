/*
* @Author: zh
* @Email: zhangheng9394@163.com
* @Date: 2020/12/11 14:35
 */
package models

type Resp map[interface{}]interface{}

type ResData struct {
	Status int
	Msg string
	Result interface{}
}