#### 一、生成可执行文件
```
# run.exe是生成的可执行exe的名称
go build -o run.exe
```

#### 二、orm使用注意问题
`使用gorm`
```
// 定义的struct中的json 名称是你查询返回，和新建或修改需要对应的字段名称
type Person struct {
	ID        uint   `json:"id"`
	FirstName string `gorm:"column:name" json:"first_name"`
	//FirstName string `json:"first_name"`
	LastName  string `json:"last_name" `
	Tel       string `json:"tel"`
	TelPhone  string `json:"tel_phone" `
	City      string `json:"city" `
}

// 如果post新建，对应的请求数据中，对应的字段名称是first_name,last_name等等，而不是FirstName
// 请求参数如下为正确
/*
{
    "first_name":"sss111111",
    "last_name": "sss333333",
    "tel": "123123123"
}
*/
```
#### 三、文件介绍
##### （1）配置文件
```
config/app.ini   配置文件
```
##### （2）配置文件加载
```
utils/ini_utiil.go   是配置文件初始化项目时加载的配置文件，目前加载的事web服务和数据库的初始化数据
```
##### （3）定义json格式数据
```
models/json_type.go  其他 struct 的类型定义为 JSON，则会进行数据转换，将接收的对象转换为 json
```