package models

import (
	"first-beego/src/first-beego/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int
	CreateTime int64
}

func InsertUser(user User) (int64, error) {
	return utils.ModifyDB(
		"insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.CreateTime)
}

func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

func QueryUserByUsername(username string) int {
	sql := fmt.Sprintf("where username = '%s'", username)
	return QueryUserWightCon(sql)
}

func QueryUserByUsernameAndPwd(username, password string) int {
	sql := fmt.Sprintf("where username = '%s' and password = '%s'", username, password)
	return QueryUserWightCon(sql)
}
