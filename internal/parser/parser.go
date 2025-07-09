package parser

import (
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

var imgRegexp = regexp.MustCompile(`<img[^>]+src="([^"]+)"`)

func ParseRSS(url string) ([]Article, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	var articles []Article
	for _, item := range feed.Items {
		image := extractFirstImageURL(item.Description)
		description := stripHTML(removeImgTags(item.Description))
		author := ""
		if item.Author != nil {
			author = item.Author.Name
		}
		published := time.Time{}
		if item.PublishedParsed != nil {
			published = *item.PublishedParsed
		}

		a := Article{
			Title:       item.Title,
			Link:        item.Link,
			Author:      author,
			Published:   published,
			Description: description,
			Image:       image,
			Categories:  item.Categories,
		}
		articles = append(articles, a)
	}

	return articles, nil
}

func extractFirstImageURL(input string) string {
	match := imgRegexp.FindStringSubmatch(input)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func removeImgTags(input string) string {
	re := regexp.MustCompile(`<img[^>]*>`)
	return re.ReplaceAllString(input, "")
}

func stripHTML(input string) string {
	inTag := false
	var result strings.Builder

	for _, r := range input {
		switch r {
		case '<':
			inTag = true
		case '>':
			inTag = false
		default:
			if !inTag {
				result.WriteRune(r)
			}
		}
	}
	return strings.TrimSpace(result.String())
}
