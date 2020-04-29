/*
   @Time : 2020/4/22 4:38 下午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : configer
   @Software: GoLand
*/

package configer

import (
	"os"
	"strings"
)

type config struct {
	Project string // 项目名
	Version string // 项目版本
	Debug   bool   // 调试模式
}

// Conf 是项目的基本配置
// 包括项目名称, 项目版本, 是否开启调试模式
// 此部分项目建议在进行打包程序时进行设置, 例如 :
// go build -ldflags "-X configer.Conf.Project=YourProject"
var Conf = config{
	Project: "Go-Common",
	Version: "0.1.0",
	Debug:   GetBool("Debug", false),
}

// GetString 用于获取字符型环境变量
func GetString(parameter string, def string) string {
	if os.Getenv(parameter) != "" {
		return os.Getenv(parameter)
	}
	return def
}

// GetBool 用于获取布尔型环境变量
func GetBool(parameter string, def bool) bool {
	if os.Getenv(parameter) != "" {
		if strings.ToLower(os.Getenv(parameter)) == "true" {
			return true
		}
		return false
	}
	return def
}
