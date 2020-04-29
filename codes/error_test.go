/*
   @Time : 2020/4/29 1:38 下午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : error_test
   @Software: GoLand
*/

package codes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorText(t *testing.T) {
	// 测试已知错误代码的文本返回是否符合预期
	assert.Equal(t, errorText[InternalErrorMissingConfig], ErrorText(InternalErrorMissingConfig))
	// 测试未知错误代码的文本返回是否符合预期
	assert.Equal(t, errorText[InternalErrorUnknown], ErrorText(-1))
}
