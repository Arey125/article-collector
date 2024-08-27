package models

import (
	"database/sql"
	"errors"

	"github.com/Arey125/article-collector/internal/article"
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
    source, ok := SourceMap[sourceId]
    if (!ok) {
        return nil, errors.New("No such source")
    }

    stmt := "SELECT name, link, status_id FROM articles WHERE source_id = ?"
    rows, err := model.DB.Query(stmt, sourceId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    articles := []Article{}

    for rows.Next() {
        article := article.NewArticle("", "", source)

        err := rows.Scan(&article.Name, &article.Link, &article.Status)
        if err != nil {
            return nil, err
        }

        articles = append(articles, article)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return articles, nil
}
