package controllers

import (
	"first-beego/src/first-beego/models"
	"first-beego/src/first-beego/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

// 处理注册
func (this *RegisterController) Post() {

	// 获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)
	logs.Info(username, password, repassword)
	
	// 注册之前先判断用户是否注册，如果已经注册，返回错误
	id := models.QueryUserByUsername(username)
	fmt.Println("id:",id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": "用户已存在",
		}
		this.ServeJSON()
		return
	}
	password = utils.MD5(password)

	user := models.User{0,username,password,0,time.Now().Unix()}

	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": "注册失败",
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": "注册成功",
		}
	}

	this.ServeJSON()
}
