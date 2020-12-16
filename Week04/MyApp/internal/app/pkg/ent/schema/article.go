package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique(),
		field.String("body").SchemaType(map[string]string{
			dialect.MySQL: "longtext",
		}),
		field.Bool("published"),
		field.Time("created_time").Default(time.Now),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
