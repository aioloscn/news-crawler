package main

import (
	"news-crawler/engine"
	"news-crawler/news/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "https://www.chinanews.com/scroll-news/news1.html",
		ParserFunc: parser.ParseNewsTitle,
	})
}
