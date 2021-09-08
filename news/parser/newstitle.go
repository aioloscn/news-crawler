package parser

import (
	"news-crawler/engine"
	"regexp"
)

var newsTitleRe = regexp.MustCompile(`<a href="//(www.chinanews.com/.+.shtml)">([^<]+)</a></div>`)

func ParseNewsTitle(contents []byte) engine.ParseResult {
	matches := newsTitleRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: "https://" + string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c)
			},
		})
	}
	return result
}
