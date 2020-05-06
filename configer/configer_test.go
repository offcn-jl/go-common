/*
   @Time : 2020/4/29 2:12 下午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : configer_test
   @Software: GoLand
*/

package configer

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIsTest(t *testing.T) {
	// 测试默认状态
	assert.False(t, IsDebug())

	// 测试开启调试
	assert.NoError(t, os.Setenv("Debug", "true"))
	assert.True(t, IsDebug())

	// 测试关闭调试
	assert.NoError(t, os.Setenv("Debug", "false"))
	assert.False(t, IsDebug())
}

func TestGetString(t *testing.T) {
	assert.NoError(t, os.Setenv("TestExistStringEnv", "ExampleString"))

	assert.Equal(t, "ExampleString", GetString("TestExistStringEnv", "NotMatchString"))
	assert.Equal(t, "NotMatchString", GetString("TestNotExistStringEnv", "NotMatchString"))
}

func TestGetBool(t *testing.T) {
	assert.NoError(t, os.Setenv("TestExistBoolTrueEnv", "true"))
	assert.NoError(t, os.Setenv("TestExistBoolFalseEnv", "false"))

	assert.True(t, GetBool("TestExistBoolTrueEnv", true))
	assert.False(t, GetBool("TestExistBoolFalseEnv", false))
	assert.False(t, GetBool("TestNotExistBoolEnv", false))
}
