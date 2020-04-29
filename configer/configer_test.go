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

func TestGetString(t *testing.T) {
	os.Setenv("TestExistStringEnv", "ExampleString")

	assert.Equal(t, "ExampleString", GetString("TestExistStringEnv", "NotMatchString"))
	assert.Equal(t, "NotMatchString", GetString("TestNotExistStringEnv", "NotMatchString"))
}

func TestGetBool(t *testing.T) {
	os.Setenv("TestExistBoolTrueEnv", "true")
	os.Setenv("TestExistBoolFalseEnv", "false")

	assert.True(t, GetBool("TestExistBoolTrueEnv", true))
	assert.False(t, GetBool("TestExistBoolFalseEnv", false))
	assert.False(t, GetBool("TestNotExistBoolEnv", false))
}
