package parser

type Repository interface {
	SaveArticles(articles []Article) error
}
