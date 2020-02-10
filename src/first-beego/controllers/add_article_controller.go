package controllers

import (
	"first-beego/src/first-beego/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get(){
	this.TplName = "write_article.html"
}

func (this *AddArticleController) Post(){
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	article := models.Article{0,title,tags,short,content,"Jason",time.Now().Unix()}
	_, err := models.AddArticle(article)
	if err == nil  {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"message": "ok",
		}
	}else{
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": "error",
		}
	}
	this.ServeJSON()
}
