package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Gomziakoff/HabrTok/config"
	"github.com/Gomziakoff/HabrTok/internal/parser"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	ticker := time.NewTicker(cfg.Interval)
	defer ticker.Stop()

	run(cfg)

	for range ticker.C {
		run(cfg)
	}
}

func run(cfg *config.ParserConfig) {
	articles, err := parser.ParseRSS(cfg.RSSUrl)
	if err != nil {
		log.Printf("Ошибка парсинга RSS: %v", err)
		return
	}

	fmt.Printf("Получено статей: %d\n", len(articles))
	for _, a := range articles {
		fmt.Println("=== Статья ===")
		fmt.Println("Заголовок:", a.Title)
		fmt.Println("Ссылка:", a.Link)
		fmt.Println("Автор:", a.Author)
		fmt.Println("Дата:", a.Published.Format(time.RFC3339))
		fmt.Println("Картинка:", a.Image)
		fmt.Println("Описание:", a.Description)
		fmt.Println("Категории:", a.Categories)
		fmt.Println()
	}
}
