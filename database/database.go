/*
   @Time : 2020/4/22 3:38 下午
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : database
   @Software: GoLand
*/

package database

import (
	"errors"
	"github.com/offcn-jl/go-common/configer"
	"github.com/offcn-jl/go-common/logger"
)

/**
 * 数据库配置
 */
type DSN struct {
	PostgreSQL string
}

func GetDSN() (d DSN) {
	if configer.GetString("POSTGRE_SQL_HOST", "") == "" {
		logger.Panic(errors.New("未配置 POSTGRE_SQL_HOST"))
	}
	if configer.GetString("POSTGRE_SQL_PORT", "") == "" {
		logger.Panic(errors.New("未配置 POSTGRE_SQL_PORT"))
	}
	if configer.GetString("POSTGRE_SQL_DB_NAME", "") == "" {
		logger.Panic(errors.New("未配置 POSTGRE_SQL_DB_NAME"))
	}
	if configer.GetString("POSTGRE_SQL_USER", "") == "" {
		logger.Panic(errors.New("未配置 POSTGRE_SQL_USER"))
	}
	if configer.GetString("POSTGRE_SQL_PASSWORD", "") == "" {
		logger.Panic(errors.New("未配置 POSTGRE_SQL_PASSWORD"))
	}
	if configer.GetString("POSTGRE_SQL_SSL_MODE", "") == "" {
		logger.Panic(errors.New("未配置 POSTGRE_SQL_SSL_MODE"))
	}
	d.PostgreSQL = "host=" + configer.GetString("POSTGRE_SQL_HOST", "") + " port=" + configer.GetString("POSTGRE_SQL_PORT", "") + " user=" + configer.GetString("POSTGRE_SQL_USER", "") + " dbname=" + configer.GetString("POSTGRE_SQL_DB_NAME", "") + " password=" + configer.GetString("POSTGRE_SQL_PASSWORD", "") + " sslmode=" + configer.GetString("POSTGRE_SQL_SSL_MODE", "")
	return
}
