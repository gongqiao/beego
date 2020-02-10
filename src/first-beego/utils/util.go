package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

var db *sql.DB

func InitMySql() {
	fmt.Println("初始化 mysql 数据库")
	driverName := beego.AppConfig.String("driverName")

	// 注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	// 数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	err := orm.RegisterDataBase("default", driverName, dbConn)

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		logs.Error("数据库连接错误。。。", err.Error())
		return
	} else {
		db = db1
	}

	logs.Info("数据库连接成功")
}

// 数据库操作
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	return count, err
}


//查询
func QueryRowDB(sql string) *sql.Row {
	logs.Info(sql)
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
