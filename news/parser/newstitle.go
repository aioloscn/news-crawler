package parser

import (
	"news-crawler/engine"
	"regexp"
)

const newsTitleRe = `<a href="//(www.chinanews.com/.+.shtml)">([^<]+)</a></div>`

func ParseNewsTitle(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(newsTitleRe)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://" + string(m[1]),
			ParserFunc: ParseProfile,
		})
	}
	return result
}
