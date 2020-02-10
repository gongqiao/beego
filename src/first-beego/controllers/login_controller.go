package controllers

import (
	"first-beego/src/first-beego/models"
	"first-beego/src/first-beego/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}


// 处理页面请求
func (this *LoginController) Get(){
	this.TplName = "login.html"
}

func (this *LoginController) Post(){

	username := this.GetString("username")
	password := this.GetString("password")

	fmt.Println("username: ", username, "password: ",password)

	id := models.QueryUserByUsernameAndPwd(username,utils.MD5(password))

	fmt.Println("user id: ", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"message": "登录成功",
		}

		this.SetSession("loginuser", username)
		// loginuser -> jason
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": "登录失败",
		}
	}
	this.ServeJSON()
}