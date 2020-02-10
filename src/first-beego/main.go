package main

import (
	_ "first-beego/src/first-beego/routers"
	"first-beego/src/first-beego/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMySql()
	beego.Run()
}

