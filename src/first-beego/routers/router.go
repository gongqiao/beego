package routers

import (
	"first-beego/src/first-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.HomeController{})

    // 注册
    beego.Router("/register", &controllers.RegisterController{})

    // 登录
    beego.Router("/login", &controllers.LoginController{})

    // 退出登陆
    beego.Router("/exit", &controllers.ExitController{})

    // 添加文章
    beego.Router("/article/add", &controllers.AddArticleController{})
    beego.Router("/article/:id", &controllers.ShowArticleController{})
}
