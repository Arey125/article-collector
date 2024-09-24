package server

import (
	"path"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/Arey125/article-collector/internal/server/template"
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

func getArticleNav(article article.Article) []Link {
	return []Link{
		{Title: "Home", Link: "/"},
		getSourceLink(*article.Source),
		getArticleLink(article),
	}
}

func getSourceNav(source article.Source) []Link {
	return []Link{
		{Title: "Home", Link: "/"},
		getSourceLink(source),
	}
}

func getSourceTemplateLink(source article.Source) template.TemplateLink {
	return template.TemplateLink{
		Title: source.Name,
		Link:  path.Join("/source", source.Domain),
	}
}

func getArticleTemplateLink(article article.Article) template.TemplateLink {
	return template.TemplateLink{
		Title: article.Name,
		Link:  path.Join("/source", article.Source.Domain, path.Base(article.Link)),
	}
}

func getArticleTemplateNav(article article.Article) []template.TemplateLink {
	return []template.TemplateLink{
		{Title: "Home", Link: "/"},
		getSourceTemplateLink(*article.Source),
		getArticleTemplateLink(article),
	}
}

func getSourceTemplateNav(source article.Source) []template.TemplateLink {
	return []template.TemplateLink{
		{Title: "Home", Link: "/"},
		getSourceTemplateLink(source),
	}
}
