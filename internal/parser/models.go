package parser

import "time"

type Article struct {
	Title       string
	Link        string
	Author      string
	Published   time.Time
	Description string
	Image       string
	Categories  []string
}
