package models

import (
	"bytes"
	"first-beego/src/first-beego/utils"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       [] TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string

	// 查看文章地址
	Link string

	//修改文章地址
	UpdateLink string
	DeleteLink string

	// 记录是否登录
	IdLogin bool
}

type TagLink struct {
	TagName, TagUrl string
}

// 分页结构体
type HomeFooterPageCode struct {
	HasPre, HasNext             bool
	ShowPage, PreLink, NextLink string
}

func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		homeParam := HomeBlockParam{
			Id:         art.Id,
			Title:      art.Title,
			Tags:       createTagsLinks(art.Tags),
			Short:      art.Short,
			Content:    art.Content,
			Author:     art.Author,
			CreateTime: utils.SwitchTimeStampToData(art.CreateTime),
			Link:       "/article/" + strconv.Itoa(art.Id),
			UpdateLink: "/article/update?id=" + strconv.Itoa(art.Id),
			DeleteLink: "/article/delete?id=" + strconv.Itoa(art.Id),
			IdLogin:    isLogin,
		}

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}

func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}

	// 查询出总条数
	num := GetArticleRowsNum()

	//从配置文件中读取每页显示多少条
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")

	// 在计算出总页数
	fmt.Println(num)
	allPageNum := (num-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}

func createTagsLinks(tags string) []TagLink {
	var tagLink [] TagLink
	tagsParam := strings.Split(tags, "&")

	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}
