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
	// 测试 未配置 POSTGRE_SQL_HOST
	assert.PanicsWithError(t, "未配置 POSTGRE_SQL_HOST", func() { GetDSN() })
	// 配置 POSTGRE_SQL_HOST
	assert.NoError(t, os.Setenv("POSTGRE_SQL_HOST", "POSTGRE_SQL_HOST"))

	// 测试 未配置 POSTGRE_SQL_PORT
	assert.PanicsWithError(t, "未配置 POSTGRE_SQL_PORT", func() { GetDSN() })
	// 配置 POSTGRE_SQL_PORT
	assert.NoError(t, os.Setenv("POSTGRE_SQL_PORT", "POSTGRE_SQL_PORT"))

	// 测试 未配置 POSTGRE_SQL_DB_NAME
	assert.PanicsWithError(t, "未配置 POSTGRE_SQL_DB_NAME", func() { GetDSN() })
	// 配置 POSTGRE_SQL_DB_NAME
	assert.NoError(t, os.Setenv("POSTGRE_SQL_DB_NAME", "POSTGRE_SQL_DB_NAME"))

	// 测试 未配置 POSTGRE_SQL_USER
	assert.PanicsWithError(t, "未配置 POSTGRE_SQL_USER", func() { GetDSN() })
	// 配置 POSTGRE_SQL_USER
	assert.NoError(t, os.Setenv("POSTGRE_SQL_USER", "POSTGRE_SQL_USER"))

	// 测试 未配置 POSTGRE_SQL_PASSWORD
	assert.PanicsWithError(t, "未配置 POSTGRE_SQL_PASSWORD", func() { GetDSN() })
	// 配置 POSTGRE_SQL_PASSWORD
	assert.NoError(t, os.Setenv("POSTGRE_SQL_PASSWORD", "POSTGRE_SQL_PASSWORD"))

	// 测试 未配置 POSTGRE_SQL_SSL_MODE
	assert.PanicsWithError(t, "未配置 POSTGRE_SQL_SSL_MODE", func() { GetDSN() })
	// 配置 POSTGRE_SQL_SSL_MODE
	assert.NoError(t, os.Setenv("POSTGRE_SQL_SSL_MODE", "POSTGRE_SQL_SSL_MODE"))

	// 测试获取 DSN
	assert.Equal(t, "host=POSTGRE_SQL_HOST port=POSTGRE_SQL_PORT user=POSTGRE_SQL_USER dbname=POSTGRE_SQL_DB_NAME password=POSTGRE_SQL_PASSWORD sslmode=POSTGRE_SQL_SSL_MODE", GetDSN().PostgreSQL)
}
