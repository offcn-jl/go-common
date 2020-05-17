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

func IsDebug() bool {
	return GetBool("DEBUG", false)
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
