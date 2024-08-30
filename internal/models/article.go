package models

import (
	"database/sql"
	"errors"
	"sort"

	. "github.com/Arey125/article-collector/internal/article"
)

type ArticleModel struct {
	DB *sql.DB
}

func (model *ArticleModel) InsertOrReplace(article *Article) error {
	stmt := `
    INSERT OR REPLACE INTO articles (name, link, source_id, slug, sort, status_id)
    VALUES (?, ?, ?, ?, ?, ?)`
	_, err := model.DB.Exec(stmt,
		article.Name, article.Link, article.Source.Id, article.Slug, article.Sort, article.Status)
	return err
}

func (model *ArticleModel) Get(sourceId string, slug string) (*Article, error) {
	const stmt = "SELECT name, link, status_id, sort FROM articles WHERE source_id = ? AND slug = ?"

	row := model.DB.QueryRow(stmt, sourceId, slug)
	if err := row.Err(); err != nil {
		return nil, err
	}

    var name, link, statusId string
    var sort int

    err := row.Scan(&name, &link, &statusId, &sort)
    if err != nil {
        return nil, err
    }

    article := NewArticle(name, link, SourceMap[sourceId])
    article.Status = statusId
    article.Sort = sort

    return &article, nil
}

func (model *ArticleModel) FromSource(sourceId string) ([]Article, error) {
	source, ok := SourceMap[sourceId]
	if !ok {
		return nil, errors.New("No such source")
	}

	stmt := "SELECT name, link, status_id, sort FROM articles WHERE source_id = ?"
	rows, err := model.DB.Query(stmt, sourceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := []Article{}

	for rows.Next() {
        var name, link, statusId string
        var sort int

		err := rows.Scan(&name, &link, &statusId, &sort)
		if err != nil {
			return nil, err
		}

		article := NewArticle(name, link, source)
        article.Status = statusId
        article.Sort = sort
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

    sort.Slice(articles, func(i, j int) bool {
        return articles[i].Sort < articles[j].Sort
    })

	return articles, nil
}
