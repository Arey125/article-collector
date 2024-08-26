package models

import (
	"database/sql"

	. "github.com/Arey125/article-collector/internal/article"
)

type ArticleModel struct {
	DB *sql.DB
}

func (model *ArticleModel) InsertOrReplace(article *Article) error {
	stmt := "INSERT OR REPLACE INTO articles (name, link, source_id) VALUES (?, ?, ?)"
	_, err := model.DB.Exec(stmt,
		article.Name, article.Link, article.Source.Id)
	return err
}

func (model *ArticleModel) FromSource(sourceId string) ([]Article, error) {
    return nil, nil // TODO
}
