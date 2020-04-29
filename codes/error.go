/*
   @Time : 2019-08-31 08:49
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : error
   @Software: GoLand
*/

package codes

// 定义错误代码
const (
	// 内部错误 10000
	InternalErrorUnknown           = 10000
	InternalErrorMissingConfig     = 10001
	InternalErrorHandler           = 10002
	InternalErrorRaw               = 10003
	InternalErrorDecryptXML        = 10004
	InternalErrorUnmarshalXML      = 10005
	InternalErrorSaveAuthorized    = 10006
	InternalErrorGetART            = 10007
	InternalErrorGetAAT            = 10008
	InternalErrorSaveToORM         = 10009
	InternalErrorRequest           = 10010
	InternalErrorInvalidResponse   = 10011
	InternalErrorSendGETRequest    = 10012
	InternalErrorSendPOSTRequest   = 10013
	InternalErrorCheckSQLInjection = 10014
	InternalErrorSendGetJson       = 10015
	InternalErrorGetMISToken       = 10016
	InternalErrorMarshalJson       = 10017
	InternalErrorUnmarshalJson     = 10018
	InternalErrorBase64Encode      = 10019
	InternalErrorBase64Decode      = 10020
	InternalErrorRSAEncryption     = 10021
	InternalErrorRSADecryption     = 10022
	InternalErrorGetContextInfo    = 10023
	InternalErrorGetWOPC           = 10024
	InternalErrorGetPAC            = 10025
	InternalErrorQueryUnescape     = 10026

	// 数据不合法 11000
	InvalidData                  = 11000
	InvalidJson                  = 11001
	InvalidTableName             = 11002
	InvalidPhoneNumber           = 11003
	InvalidQuery                 = 11004
	InvalidParameter             = 11005
	InvalidParameterSQLInjection = 11006

	// 数据不存在
	NotExist           = 12000
	NotExistTable      = 12001
	NotExistPage       = 12002
	NotExistSwiper     = 12003
	NotExistClass      = 12004
	NotExistAuthorizer = 12005

	// 授权错误
	NotCertified               = 13000
	NotCertifiedCORS           = 13001
	NotCertifiedUserName       = 13002
	NotCertifiedPassword       = 13003
	NotCertifiedLoginFailCount = 13004
	NotCertifiedSession        = 13005
	NotCertifiedMISToken       = 13006
	NotCertifiedPermission     = 13007

	// 操作错误
	WrongOperation            = 14000
	HasChildGroups            = 14001
	HasUsers                  = 14002
	DuplicateUserName         = 14003
	AlreadyExistPage          = 14004
	SwiperExceedTheUpperLimit = 14005
	SmsCodeNotExpired         = 14006
	SmsCodeExpired            = 14007
	SmsCodeError              = 14008
	AlreadyExistGroup         = 14009
)

// 定义错误代码对应的文本
var errorText = map[int]string{
	// 内部错误
	InternalErrorUnknown:           "Unknown Internal Error",
	InternalErrorMissingConfig:     "System Config Is Missing",
	InternalErrorHandler:           "Handler Return Error",
	InternalErrorRaw:               "Get Raw Data Fail",
	InternalErrorDecryptXML:        "Decrypt XML Fail",
	InternalErrorUnmarshalXML:      "Unmarshal XML Fail",
	InternalErrorSaveAuthorized:    "Save Authorized Info Fail",
	InternalErrorGetART:            "Get Authorizer Refresh Token Fail",
	InternalErrorGetAAT:            "Get Authorizer Access Token Fail",
	InternalErrorSaveToORM:         "Save To ORM Fail",
	InternalErrorRequest:           "Send Request Fail",
	InternalErrorInvalidResponse:   "Got Invalid Response",
	InternalErrorSendGETRequest:    "Send GET Request Fail",
	InternalErrorSendPOSTRequest:   "Send POST Request Fail",
	InternalErrorCheckSQLInjection: "Check SQL Injection Fail",
	InternalErrorSendGetJson:       "Send Get Json Request Fail",
	InternalErrorGetMISToken:       "Get MIS Token Fail",
	InternalErrorMarshalJson:       "Marshal Json Map To Json String Fail",
	InternalErrorUnmarshalJson:     "Unmarshal Json String To Json Map Fail",
	InternalErrorBase64Encode:      "Encode String To Base64 String fail",
	InternalErrorBase64Decode:      "Decode String To Base64 String fail",
	InternalErrorRSAEncryption:     "Encryption String To RSA String fail",
	InternalErrorRSADecryption:     "Decryption String To RSA String fail",
	InternalErrorGetContextInfo:    "Get Info From Context Fail",
	InternalErrorGetWOPC:           "Get Wechat Open Platform Config Fail",
	InternalErrorGetPAC:            "Get Pre Auth Code Fail",
	InternalErrorQueryUnescape:     "Unescape Query ( IN URL ) Fail",

	// 数据不合法
	InvalidData:                  "Invalid Data",
	InvalidJson:                  "Invalid Json Data",
	InvalidTableName:             "Invalid Table Name",
	InvalidPhoneNumber:           "Invalid Phone Number",
	InvalidQuery:                 "Invalid Query Url",
	InvalidParameter:             "Invalid Parameter",
	InvalidParameterSQLInjection: "Did Not Pass The SQL Injection Check",

	// 数据不存在
	NotExist:           "Data Is Not Exist",
	NotExistTable:      "Table Is Not Exist",
	NotExistPage:       "Page Is Not Exist",
	NotExistSwiper:     "Swiper Is Not Exist",
	NotExistClass:      "Class Is Not Exist",
	NotExistAuthorizer: "Authorizer Is Not Exist",

	// 授权错误
	NotCertified:               "Not Certified",
	NotCertifiedCORS:           "Not Certified ( CORS )",
	NotCertifiedUserName:       "User Has Been Banned Or Not Exist",
	NotCertifiedPassword:       "Password Not Match",
	NotCertifiedLoginFailCount: "User Has Been Login Failed More Than 5 Times In 24 Hour",
	NotCertifiedSession:        "Session Has Been Expired",
	NotCertifiedMISToken:       "MIS Token Has Been Expired",
	NotCertifiedPermission:     "User Does Not Have Permission For This Operation",

	// 操作错误
	WrongOperation:            "Wrong Operation",
	HasChildGroups:            "Target Group Has Child Group ( s )",
	HasUsers:                  "Target Group Has Users ( s )",
	DuplicateUserName:         "Duplicate User Name",
	AlreadyExistPage:          "Already Exist Page For This Group",
	SwiperExceedTheUpperLimit: "Swiper Exceed The Upper Limit",
	SmsCodeNotExpired:         "Last Sms Code Was Not Expired",
	SmsCodeExpired:            "Sms Code Was Expired",
	SmsCodeError:              "Sms Code Was Mismatch Last One",
	AlreadyExistGroup:         "Already Exist This Group",
}

// ErrorText 用于获取错误代码对应的错误文本
func ErrorText(code int) string {
	return errorText[code]
}
