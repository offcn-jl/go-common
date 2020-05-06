/*
   @Time : 2020/5/4 5:26 下午
   @Author : ShadowWalker
   @Email : master@rebeta.cn
   @File : database_test
   @Software: GoLand
*/

package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetDSN(t *testing.T) {
	// 测试 未配置 PostgreSQLHost
	assert.PanicsWithError(t, "未配置 PostgreSQLHost", func() { GetDSN() })
	// 配置 PostgreSQLHost
	assert.NoError(t, os.Setenv("PostgreSQLHost", "PostgreSQLHost"))

	// 测试 未配置 PostgreSQLPort
	assert.PanicsWithError(t, "未配置 PostgreSQLPort", func() { GetDSN() })
	// 配置 PostgreSQLHost
	assert.NoError(t, os.Setenv("PostgreSQLPort", "PostgreSQLPort"))

	// 测试 未配置 PostgreSQLPort
	assert.PanicsWithError(t, "未配置 PostgreSQLUser", func() { GetDSN() })
	// 配置 PostgreSQLHost
	assert.NoError(t, os.Setenv("PostgreSQLUser", "PostgreSQLUser"))

	// 测试 未配置 PostgreSQLPort
	assert.PanicsWithError(t, "未配置 PostgreSQLDBName", func() { GetDSN() })
	// 配置 PostgreSQLHost
	assert.NoError(t, os.Setenv("PostgreSQLDBName", "PostgreSQLDBName"))

	// 测试 未配置 PostgreSQLPort
	assert.PanicsWithError(t, "未配置 PostgreSQLPassword", func() { GetDSN() })
	// 配置 PostgreSQLHost
	assert.NoError(t, os.Setenv("PostgreSQLPassword", "PostgreSQLPassword"))

	// 测试 未配置 PostgreSQLPort
	assert.PanicsWithError(t, "未配置 PostgreSQLSSLMode", func() { GetDSN() })
	// 配置 PostgreSQLSSLMode
	assert.NoError(t, os.Setenv("PostgreSQLSSLMode", "PostgreSQLSSLMode"))

	// 测试获取 DSN
	assert.Equal(t, "host=PostgreSQLHost port=PostgreSQLPort user=PostgreSQLUser dbname=PostgreSQLDBName password=PostgreSQLPassword sslmode=PostgreSQLSSLMode", GetDSN().PostgreSQL)
}
