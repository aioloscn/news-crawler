package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCategoryList(t *testing.T) {
	contents, err := ioutil.ReadFile("categorylist_test_data.html")
	if err != nil {
		panic(err)
	}

	ParseCategoryList(contents)
}
