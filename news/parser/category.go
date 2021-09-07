package parser

import (
	"news-crawler/engine"
	"regexp"
)

const category1Re = `<a href="//(finance.ifeng.com/c/89JmTElEKZz)" target="_blank" class="imgBox-UCDWr6RD" data-innermask="true">
<img src="//(d.ifengimg.com/w144_h96_q70/x0.ifengimg.com/ucms/2021_37/669B7DE4D1250677E778C05CC3B85A2E3DAEFDF9_size80_w650_h366.jpg)" width="144" height="96" class="trans-Ba-0WPD1"></a>`

func ParseCategory(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(category1Re)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://" + string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
