/*
   @Time : 2020/4/29 2:15 下午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : logger_test
   @Software: GoLand
*/

package logger

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	// 定义接收输出的 buffer
	buffer := new(bytes.Buffer)
	// 将默认的输出 writer 修改为 buffer
	DefaultWriter = buffer
	// 输出预期字符串到日志
	Log("ExpectedString")

	assert.Contains(t, buffer.String(), "LOG")
	assert.Contains(t, buffer.String(), "ExpectedString")
}

func TestError(t *testing.T) {
	// 定义接收输出的 buffer
	buffer := new(bytes.Buffer)
	// 将默认的输出 writer 修改为 buffer
	DefaultWriter = buffer
	// 输出预期字符串到错误日志
	Error(errors.New("ExpectedErrorString"))

	assert.Contains(t, buffer.String(), "ERROR")
	assert.Contains(t, buffer.String(), "ERROR-STACKS")
	assert.Contains(t, buffer.String(), "ExpectedErrorString")
}

func TestPanic(t *testing.T) {
	// 定义接收输出的 buffer
	buffer := new(bytes.Buffer)
	// 将默认的输出 writer 修改为 buffer
	DefaultWriter = buffer
	// 输出预期字符串到 PANIC 日志
	assert.PanicsWithError(t, "ExpectedPanicString", func() { Panic(errors.New("ExpectedPanicString")) })

	assert.Contains(t, buffer.String(), "PANIC")
	assert.Contains(t, buffer.String(), "PANIC-STACKS")
	assert.Contains(t, buffer.String(), "ExpectedPanicString")
}

func TestDebugToJson(t *testing.T) {
	// 定义接收输出的 buffer
	buffer := new(bytes.Buffer)
	// 将默认的输出 writer 修改为 buffer
	DefaultWriter = buffer
	// 定义要输出的结构
	data := map[string]interface{}{"data": "value", "foo": "bar"}

	// 测试未打开调试模式的情况
	// 输出调试内容到日志
	DebugToJson("DATA", data)
	assert.NotContains(t, buffer.String(), "DATA")
	assert.NotContains(t, buffer.String(), "DEBUG")
	assert.NotContains(t, buffer.String(), "JSON")
	assert.NotContains(t, buffer.String(), "{\"data\":\"value\",\"foo\":\"bar\"}")

	// 测试打开调试模式的情况
	// 打开调试模式
	assert.NoError(t, os.Setenv("Debug", "true"))
	// 输出调试内容到日志
	DebugToJson("DATA", data)
	assert.Contains(t, buffer.String(), "DATA")
	assert.Contains(t, buffer.String(), "DEBUG")
	assert.Contains(t, buffer.String(), "JSON")
	assert.Contains(t, buffer.String(), "{\"data\":\"value\",\"foo\":\"bar\"}")
}

func TestDebugToString(t *testing.T) {
	// 定义接收输出的 buffer
	buffer := new(bytes.Buffer)
	// 将默认的输出 writer 修改为 buffer
	DefaultWriter = buffer
	// 定义要输出的结构
	data := map[string]interface{}{"data": "value", "foo": "bar"}

	// 测试未打开调试模式的情况
	// 输出调试内容到日志
	DebugToJson("DATA", data)
	assert.NotContains(t, buffer.String(), "DATA")
	assert.NotContains(t, buffer.String(), "DEBUG")
	assert.NotContains(t, buffer.String(), "STRING")
	assert.NotContains(t, buffer.String(), "map[data:value foo:bar]")

	// 打开调试模式
	assert.NoError(t, os.Setenv("Debug", "true"))
	// 输出调试内容到日志
	DebugToString("DATA", data)
	assert.Contains(t, buffer.String(), "DATA")
	assert.Contains(t, buffer.String(), "DEBUG")
	assert.Contains(t, buffer.String(), "STRING")
	assert.Contains(t, buffer.String(), "map[data:value foo:bar]")
}
