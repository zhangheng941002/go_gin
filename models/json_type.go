/*
 * @Author: Zh
 * @Email: zhangheng9394@163.com
 * @Project: gin_web
 * @Package: models
 * @File: json_type.go
 * @Date: 2020/12/24 14:43
 */

// 其他 struct 的类型定义为下面的 JSON，则会进行数据转换，将接收的对象转换为 json

package models

import (
	"bytes"
	"database/sql/driver"
	"errors"
	jsoniter "github.com/json-iterator/go"
)

type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[:], s...)
	return nil
}
func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[:], data...)
	return nil
}

func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

func StructToJSON(value interface{}) JSON {
	bts, _ := jsoniter.Marshal(value)
	return bts
}
