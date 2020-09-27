package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Article struct{
	gorm.Model
	db *gorm.DB
	title string
	content string
	ext string
}


/**
* 初始化db
 */
func (a *Article) _init() (err error){
	a.db, err = gorm.Open("sqlite3", "article.db")
	if nil != err{
		panic(fmt.Sprintf("db[%s] connect failed!", "test.db"))
	}
	defer a.db.Close()
	a.db.AutoMigrate(&Article{})
	return nil
}

func (a *Article) Create() (err error){
	// 判断db是否链接
	if nil == a.db{
		err = a._init()
		if nil != err{
			return
		}
	}
	// 创建数据库
	a.db.Create(&Article{title : "title",content: "content",ext: "{\"summary\":\"summary content\"}"})
	return nil
}

func (a *Article) getList() (ctx *gin.Context, err error){
	// 判断db是否链接
	if nil == a.db{
		err = a._init()
		if nil != err{
			return
		}
	}
	// 查询
	var article Article
	a.db.First(&article, 1)
	//a.db.First(&article, "title = ?", "title")
	param := gin.Param{"title", article.title}
	ctx.Params = append(ctx.Params, param)
	return
}

