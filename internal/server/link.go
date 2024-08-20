package server

import (
	"path"

	"github.com/Arey125/article-collector/internal/article"
)

type Link struct {
	Title string
	Link  string
}

func getSourceLink(source article.Source) Link {
    return Link{
        Title: source.Name,
        Link:  path.Join("/source", source.Domain),
    }
}

func getArticleLink(article article.Article) Link {
    return Link{
        Title: article.Name,
        Link:  path.Join("/source", article.Source.Domain, path.Base(article.Link)),
    }
}
