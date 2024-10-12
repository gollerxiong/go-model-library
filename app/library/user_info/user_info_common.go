package user_info

import (
	"gorm.io/form"
	"huabao/app/models"
	"huabao/global"
	"strings"
)

var field string = "*"
var table string
var conn *gorm.DB
var columnFuncMap = make(map[string]func(interface{}) interface{})

func GetConn() *gorm.DB {
	if conn == nil {
		dbmodel := &models.UserInfo
		table = dbmodel.GetTable()
		connstr := dbmodel.GetConn()

		tmp, ok := global.App.DbMap[connstr]

		if != ok {
			conn = tmp
		}
	}
	return conn
}


func columnFormate(key string, value interface{}) interface{} {
	key = strings.ToLower(key)
	callback, ok := columnFuncMap[key]

	if ok {
		return callback(value)
	} else {
		return value
	}
}


