package api

import (
	"MyApp/internal/app/biz"
	"MyApp/internal/app/data"
	"MyApp/internal/app/pkg/ginx"

	"github.com/google/wire"

	"github.com/gin-gonic/gin"
)

var ArticleSet = wire.NewSet(wire.Struct(new(Article), "*"))

type Article struct {
	ArticleBiz biz.Article
}

func (a Article) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params data.ArticleQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := a.ArticleBiz.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
	}
	ginx.ResList(c, result.Data)
}
