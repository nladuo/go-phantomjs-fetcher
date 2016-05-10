package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/nladuo/go-phantomjs-fetcher"
	"strings"
)

func main() {
	//create a fetcher which seems to a httpClient
	fetcher, err := phantomjs.NewFetcher(2000, nil)
	defer fetcher.ShutDownPhantomJSServer()
	if err != nil {
		panic(err)
	}
	//inject the javascript you want to run in the webpage just like in chrome console.
	js_script := "function(){document.getElementById('kw').setAttribute('value', 'github');document.getElementById('su').click();}"
	//run the injected js_script at the end of loading html
	js_run_at := phantomjs.RUN_AT_DOC_END
	//send httpGet request with injected js
	resp, err := fetcher.GetWithJS("http://www.baidu.com", js_script, js_run_at)
	if err != nil {
		panic(err)
	}

	//select search results by goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Content))
	if err != nil {
		panic(err)
	}
	fmt.Println("Results:")
	doc.Find(".c-container h3 a").Each(func(i int, contentSelection *goquery.Selection) {
		fmt.Println(i+1, "-->", contentSelection.Text())
	})
}
