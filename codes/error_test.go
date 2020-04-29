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
	assert.Equal(t, errorText[InternalErrorUnknown], ErrorText(InternalErrorUnknown))
}
