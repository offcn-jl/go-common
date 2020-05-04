/*
   @Time : 2020/4/19 9:36 上午
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : logger
   @Software: GoLand
*/

package logger

import (
	"encoding/json"
	"fmt"
	"github.com/offcn-jl/go-common/configer"
	"io"
	"os"
	"runtime/debug"
	"time"
)

// DefaultWriter 用于日志输出的 io.Writer
var DefaultWriter io.Writer = os.Stdout

// Log 用于打印日志
func Log(log string) {
	_, _ = fmt.Fprintf(DefaultWriter, "[LOG] %v | %s\n",
		time.Now().Format("2006/01/02 - 15:04:05"),
		log,
	)
}

// Error 用于打印错误
func Error(err error) {
	_, _ = fmt.Fprintf(DefaultWriter, "[ERROR] %v | %s\n",
		time.Now().Format("2006/01/02 - 15:04:05"),
		err,
	)
	_, _ = fmt.Fprintf(DefaultWriter, "[ERROR-STACKS] %v\n%s\n",
		time.Now().Format("2006/01/02 - 15:04:05"),
		debug.Stack(), // 输出调用堆栈
	)
	// debug.PrintStack() // 打印调用堆栈
}

// Panic 用于打印错误后抛出 panic
func Panic(err error) {
	_, _ = fmt.Fprintf(DefaultWriter, "[PANIC] %v | %s\n",
		time.Now().Format("2006/01/02 - 15:04:05"),
		err,
	)
	_, _ = fmt.Fprintf(DefaultWriter, "[PANIC-STACKS] %v\n%s\n",
		time.Now().Format("2006/01/02 - 15:04:05"),
		debug.Stack(), // 输出调用堆栈
	)
	panic(err)
}

// DebugToJson 用于在调试模式开启时输出 Json 字符串
func DebugToJson(name string, parameters interface{}) {
	if configer.IsDebug() {
		jsonStrings, _ := json.Marshal(parameters)
		_, _ = fmt.Fprintf(DefaultWriter, "[DEBUG-JSON] %v | %s --> %s\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			name,
			jsonStrings,
		)
	}
}

// DebugToString 用于在调试模式开启时输出字符串
func DebugToString(name string, str interface{}) {
	if configer.IsDebug() {
		_, _ = fmt.Fprintf(DefaultWriter, "[DEBUG-STRING] %v | %s --> %s\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			name,
			str,
		)
	}
}
