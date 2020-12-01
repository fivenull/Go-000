/*
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var DB *sqlx.DB

// dao 层
type Article struct {
	ID      int
	Title   string
	Content string
}

func GetArticleList() (articleList []*Article, err error) {
	sqlStr := `select id, title, content from article`
	selectError := DB.Select(&articleList, sqlStr)
	if selectError != nil {
		// 对于dao层出现错误，这里选择通过errors.Wrapf包一下
		err = errors.Wrapf(selectError, "db select article error, sql str is:\n%s", sqlStr)
	}
	return
}

func initDB(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return errors.Wrapf(err, "sqlx open fail")
	}
	err = DB.Ping()
	if err != nil {
		return errors.Wrapf(err, "DB ping fail")
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

// 业务层
func ArticleList(c *gin.Context) {
	articleList, err := GetArticleList()
	if err != nil {
		// 业务层这里进行日志的打印，并打印出堆栈信息
		fmt.Printf("article list error:%+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": articleList,
	})
}

func main() {
	router := gin.Default()
	dns := "root:123456@tcp(192.168.1.100:3306)/Blog?parseTime=true"
	err := initDB(dns)
	if err != nil {
		panic(err)
	}
	router.GET("/article/list", ArticleList)
	router.Run(":7878")
}
