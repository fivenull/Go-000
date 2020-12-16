package router

import "github.com/gin-gonic/gin"

func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	g.Use(gin.Recovery())

	v1 := g.Group("v1")
	{
		gArticle := v1.Group("articles")
		{
			gArticle.GET("", a.ArticleAPI.Query)
		}
	}
}
