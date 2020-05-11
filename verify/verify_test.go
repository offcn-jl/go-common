/*
   @Time : 2020/5/11 1:36 下午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : verify_test
   @Software: GoLand
*/

package verify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhone(t *testing.T) {
	assert.True(t, Phone("17866668888"))      // 178 号段
	assert.True(t, Phone("17166668888"))      // 171 号段
	assert.True(t, Phone("+8617166668888"))   // 带国家码 +86
	assert.True(t, Phone("008617166668888"))  // 带国家码 0086
	assert.False(t, Phone("27866668888"))     // 287 号段
	assert.False(t, Phone("+0117866668888"))  // 带国家码 +01
	assert.False(t, Phone("000117866668888")) // 带国家码 0001
}
