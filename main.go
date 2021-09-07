package main

import (
	"news-crawler/engine"
	"news-crawler/news/parser"
	"news-crawler/persist"
	"news-crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 20,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "https://www.chinanews.com/scroll-news/news1.html",
		ParserFunc: parser.ParseNewsTitle,
	})
}
