package biz

import (
	"MyApp/internal/app/data"
	"context"

	"github.com/google/wire"
)

var ArticleSet = wire.NewSet(wire.Struct(new(Article), "*"))

type Article struct {
	ModelArticle *data.ModelArticle
}

func (a *Article) Query(ctx context.Context, params data.ArticleQueryParam) (data.ArticleQueryResult, error) {
	return a.ModelArticle.Query(ctx, params)
}
