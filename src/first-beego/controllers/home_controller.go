package controllers

import "first-beego/src/first-beego/models"

type HomeController struct{
	// beego.Controller
	 BaseController
}

func (this *HomeController) Get(){
	/*this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "home.html"*/

	tag := this.GetString("tag")
	page, _ := this.GetInt("page")

	var artList []models.Article
	if len(tag) > 0 {
		artList, _ = models.QueryArticleByTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}

		// 设置分页
		artList, _ = models.FindArticleByPage(page)

		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	this.Data["Content"] = models.MakeHomeBlocks(artList,this.IsLogin)

	this.TplName = "home.html"

}
