package parser

import (
	"news-crawler/engine"
	"regexp"
)

const categoryListRe = `<a href="https://([a-z]+.163.com/)">([^<]+)</a>`

func ParseCategoryList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(categoryListRe)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://" + string(m[1]),
			ParserFunc: ParseCategory,
		})
	}
	return result
}
