package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var DB *sqlx.DB

func InitDB() *sqlx.DB {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	db := sqlx.MustConnect("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local")
	db.SetMaxIdleConns(5)  // 最大空闲连接数
	db.SetMaxOpenConns(10) // 最大连接数

	DB = db
	return db
}

func GetDB() *sqlx.DB {
	return DB
}
