package parser

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"news-crawler/engine"
	"news-crawler/model"
	"regexp"
	"strings"
)

var idRe = regexp.MustCompile(`<div id="SOHUCS" sid="([0-9]+)"></div>`)
var categoryRe = regexp.MustCompile(`<h4 class='newsRecommendTitle'>(.+)精选：</h4>`)
var pubtimeRe = regexp.MustCompile(`<span id="pubtime_baidu">([^<]+)</span>`)
var titleRe = regexp.MustCompile(`<h1 style="display:block; position:relative; clear:both">([^<]+)</h1>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	profile.Category = string([]rune(extractString(contents, categoryRe))[:2])
	profile.Pubtime = extractString(contents, pubtimeRe)
	profile.Title = extractString(contents, titleRe)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find(".left_zw").Each(func(i int, selection *goquery.Selection) {
		profile.Content, _ = selection.Html()
	})

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Id:      extractString(contents, idRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
