# 代码库

系统中的代码均在本库中进行定义，规范如下：

**0. 原则上不允许对已经定义在库中的代码进行变更**

1. 代码库中需要分别定义代码及代码对应的文本内容

2. 代码的名称应当是的简称

3. 代码的类型应当是整型 ( int ) 的常量 ( const )

4. 代码文本应当是下标为整型的一维数组

5. 代码文本的数组下标应当直接引用代码常量

6. 代码文本数组中的元素内容应当是的详细介绍

7. 应当提供获取代码文本的函数，参数为代码

8. 使用代码时，应当使用代码常量名，而不是直接使用代码

9. 使用代码文本时，应当使用函数获取，而不是直接使用文本

采用如上规范的好处是：

a. 可以保证代码在系统内是唯一的

b. 可以保证代码不会在不同的模块中产生歧义

c. 如果代码的数值发生了变更，只需要更新代码库，即保证系统中所有模块都进行对应的变更

example:

```
// 代码
const (
	ExampleCodeA = 10000
	ExampleCodeB = 10001
)
// 代码文本
var codeText = map[int]string{
	ExampleCodeA: "代码 A 代表的含义",
	ExampleCodeB: "代码 A 代表的含义",
}
// 获取代码文本的函数
func CodeText(code int) string {
	return codeText[code]
}
```

错误代码表：

| 代码 | 文本描述 | 附加内容 |
|  :----:  | :----:  | :----:  |
| 10xxx | 系统内部错误 | - |
| 10000 | Unknown Internal Error | - |
| 10001 | System Config Is Missing | Item ( String ) |
| 10002 | Handler Return Error | - |
| 10003 | Get Raw Data Fail | - |
| 10004 | Decrypt XML Fail | - |
| 10005 | Unmarshal XML Fail | Parameter ( String ) |
| 10006 | Save Authorized Info Fail | Parameter ( String ) |
| 10007 | Get Authorizer Refresh Token Fail | - |
| 10008 | Get Authorizer Access Token Fail | - |
| 10009 | Save To ORM Fail | - |
| 10010 | Send Request Fail | - |
| 10011 | Got Invalid Response | Action ( String ) |
| 10012 | Send GET Request Fail | - |
| 10013 | Send POST Request Fail | - |
| 10014 | Check SQL Injection Fail | - |
| 10015 | Send Get Json Request Fail | - |
| 10016 | Get MIS Token Fail | Status, Message |
| 10017 | Marshal Json Map To Json String Fail | Parameter ( String ) |
| 10018 | Unmarshal Json String To Json Map Fail | Parameter ( String ) |
| 10019 | Encode String To Base64 String fail | Parameter ( String ) |
| 10020 | Decode String To Base64 String fail | Parameter ( String ) |
| 10021 | Encryption String To RSA String fail | Parameter ( String ) |
| 10022 | Decryption String To RSA String fail | Parameter ( String ) |
| 10023 | Get Session Info From Context Fail | Stage ( String ), Info ( String ) |
| 10024 | Get Wechat Open Platform Config Fail | - |
| 10025 | Get Pre Auth Code Fail | - |
| 10026 | Unescape Query ( IN URL ) Fail | - |
| 11xxx | 数据不合法 | - |
| 11000 | Invalid Data | Type ( String ) |
| 11001 | Invalid Json Data | - |
| 11002 | Invalid Table Name | - |
| 11003 | Invalid Phone Number | - |
| 11004 | Invalid Query Url | - |
| 11005 | Invalid Parameter | Parameter ( String ) |
| 11006 | Did Not Pass The SQL Injection Check | Parameter ( String ) |
| 12xxx | 数据不存在 | - |
| 12000 | Data Is Not Exist | Data ( String )|
| 12001 | Table Is Not Exist | TableName ( String ) |
| 12002 | Page Is Not Exist | - |
| 12003 | Swiper Is Not Exist | - |
| 12004 | Class Is Not Exist | - |
| 13xxx | 授权错误 | - |
| 13000 | Not Certified | - |
| 13001 | Not Certified ( CORS ) | - |
| 13002 | User Has Been Banned Or Not Exist | - |
| 13003 | Password Not Match | FailCount ( Int ) |
| 13004 | User Has Been Login Failed More Than 5 Times In 24 Hour | - |
| 13005 | Session Has Been Expired | - |
| 13006 | MIS Token Has Been Expired | - |
| 13007 | User Does Not Have Permission For This Operation | Details |
| 14xxx | 操作错误 | - |
| 14000 | Wrong Operation | - |
| 14001 | Target Group has Child Group ( s ) | - |
| 14002 | Target Group has Users ( s ) | - |
| 14003 | Duplicate User Name | - |
| 14004 | Already Exist Page For This Group | - |
| 14005 | Swiper Exceed The Upper Limit | - |
| 14006 | Last Sms Code Was Not Expired | - |
| 14007 | Sms Code Was Expired | - |
| 14008 | Sms Code Was Mismatch Last One | - |
| 14009 | Already Exist This Group | - |
