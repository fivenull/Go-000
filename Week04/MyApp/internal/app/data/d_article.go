package data

import (
	"MyApp/internal/app/pkg/ent"
	"MyApp/internal/app/pkg/ent/article"
	"context"
	"time"

	"github.com/google/wire"
)

var ModelArticleSet = wire.NewSet(wire.Struct(new(ModelArticle), "*"))

type ListResult struct {
	List interface{}
}

type ErrorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResult 响应错误
type ErrorResult struct {
	Error ErrorItem `json:"error"` // 错误项
}

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleQueryParam struct {
	Title string `form:"title"`
}

type ArticleQueryResult struct {
	Data Articles
}

type Articles []*Article

type ModelArticle struct {
	DB *ent.Client
}

func (m *ModelArticle) Query(ctx context.Context, params ArticleQueryParam) (ArticleQueryResult, error) {
	articleQuery := m.DB.Article.Query()
	if params.Title != "" {
		articleQuery = articleQuery.Where(article.TitleEQ(params.Title))
	}
	articles, err := articleQuery.All(ctx)
	var articleResult []*Article
	for _, article := range articles {
		articleResult = append(articleResult, &Article{
			ID:        article.ID,
			Title:     article.Title,
			Body:      article.Body,
			CreatedAt: article.CreatedTime,
		})
	}
	return ArticleQueryResult{Data: articleResult}, err
}
