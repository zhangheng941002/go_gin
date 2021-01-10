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

#### 四、swagger
`路由：/swagger/index.html` 

`需要在每个路由处理函数上加入备注才会展示出来`
##### （1）安装swag cli
`按照好后需要配置到环境变量中，通过下面命令查看是否成功：swag --version`
```
go get -u github.com/swaggo/swag/cmd/swag
```
##### （2）安装gin与swagger相关到插件
```
# gin-swagger 中间件
go get github.com/swaggo/gin-swagger

# swagger 内置文件
go get github.com/swaggo/gin-swagger/swaggerFiles
```
##### （3）相关配置
###### 项目配置

`官方demo：`https://github.com/swaggo/swag/blob/master/example/celler/main.go

`官方介绍：`https://swaggo.github.io/swaggo.io/declarative_comments_format/general_api_info.html
```
// @title 这里写标题
// @version 版本
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url 联系人的网站
// @contact.email 联系人邮箱

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
```

###### 接口配置说明
`官方介绍：`https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
```
// @Summary 就有名（该接口是提供的什么功能）
// @Description 对接口的详细描述
// @Id 是一个全局标识符，所有的接口文档中 Id 不能标注
// @Tags 是对接口的标注，同一个 tag 为一组，这样方便我们整理接口
// @Version 表明该接口的版本
// @Accept 表示该该请求的请求类型
// @Param 表示参数 分别有以下参数 参数名词 参数类型 数据类型 是否必须 注释 属性(可选参数),参数之间用空格隔开。
// @Success 表示请求成功后返回，它有以下参数 请求返回状态码，参数类型，数据类型，注释
// @Failure 请求失败后返回，参数同上
// @Router 该函数定义了请求路由并且包含路由的请求方式。
```
```
Description
描述，支持Markdown，但是是代码格式的
表格如下：
// @Description | 项目 | 价格 | 数量 |
// @Description | :-------- | --------:| :--: |
// @Description | iPhone | 6000 元 | 5 |
// @Description | iPad | 3800 元 | 12 |
// @Description | iMac | 10000 元 | 234 |

Param
参数，从前往后分别是：

@Param who query string true “人名”

@Param 1.参数名 2.参数类型 3.参数数据类型 4.是否必须 5.参数描述 6.其他属性

1.参数名

参数名就是我们解释参数的名字。

2.参数类型

参数类型主要有四种：

path 该类型参数直接拼接在URL中，如Demo中HandleGetFile：

// @Param id path integer true "文件ID"
query 该类型参数一般是组合在URL中的，如Demo中HandleHello

// @Param who query string true "人名"
formData 该类型参数一般是POST,PUT方法所用，如Demo中HandleLogin

// @Param user formData string true "用户名" default(admin)
body 当Accept是JSON格式时，我们使用该字段指定接收的JSON类型

// @Param param body main.JSONParams true "需要上传的JSON"

3.参数数据类型

数据类型主要支持一下几种：

string (string)
integer (int, uint, uint32, uint64)
number (float32)
boolean (bool)
注意，如果你是上传文件可以使用file, 但参数类型一定是formData, 如下：

// @Param file formData file true "文件"
4.是否是必须

表明该参数是否是必须需要的，必须的在文档中会黑体标出，测试时必须填写。

5.参数描述

就是参数的一些说明

6.其他属性

除了上面这些属性外，我们还可以为该参数填写一些额外的属性，如枚举，默认值，值范围等。如下：

枚举
// @Param enumstring query string false "string enums" Enums(A, B, C)
// @Param enumint query int false "int enums" Enums(1, 2, 3)
// @Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)

值添加范围
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Param int query int false "int valid" mininum(1) maxinum(10)

设置默认值
// @Param default query string false "string default" default(A)
而且这些参数是可以组合使用的，如：

// @Param enumstring query string false "string enums" Enums(A, B, C) default(A)
Success
指定成功响应的数据。格式为：

// @Success 1.HTTP响应码 {2.响应参数类型} 3.响应数据类型 4.其他描述

1.HTTP响应码

也就是200，400，500那些。

2.响应参数类型 / 3.响应数据类型

返回的数据类型，可以是自定义类型，可以是json。

自定义类型
在平常的使用中，我都会返回一些指定的模型序列化JSON的数据，这时，就可以这么写：

// @Success 200 {object} main.File
其中，模型直接用包名.模型即可。你会说，假如我返回模型数组怎么办？这时你可以这么写：

// @Success 200 {anrry} main.File
json
将如你只是返回其他的数据格式可如下写：

// @Success 200 {string} string ""
4.其他描述

可以添加一些说明。

Failure
​ 同Success。
```
`备注：Router 中如果是在URL中获取参数，需要用{}包裹。 例如：/get/all/{urlName}/test`

###### 使用
```
//在main函数的同级目录下执行以下指令，会生成docs文件夹
swag init
```