package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/refusetofeel/go-phantomjs-fetcher"
	"strings"
)

func main() {
	//create a fetcher which simulates an httpClient
	//Third parameter can be empty "", or type in a REGEX without //g. This blocks requests that are not needed to continue your scrape, in turn making it faster to load.
	//In this example, searspartsdirect.com loads .png,.jpg,.css,.woff files, these are not needed, so lets add these. You can find what you need to block by simply going too the webpage and viewing network resources.
	fetcher, err := phantomjs.NewFetcher(2016, nil,".css|.png|.jpg|.woff") //this is just a sample, can add more, or use less for the regex. This can be any normal regex, not just | seperated strings.
	defer fetcher.ShutDownPhantomJSServer()
	if err != nil {
		panic(err)
	}
	//inject the javascript you want to run in the webpage just like in chrome console.
	js_script := "function(){}"
	//run the injected js_script at the end of loading html
	js_run_at := phantomjs.RUN_AT_DOC_END
	//send httpGet request with injected js
	resp, err := fetcher.GetWithJS("http://www.searspartsdirect.com/model/search.html?q=FFTR1814QW", js_script, js_run_at)
	if err != nil {
		panic(err)
	}

	TotalTime := fmt.Sprintf("Total Response Time: %v", resp.Time)
	fmt.Println(TotalTime)

	//select search results by goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Content))
	if err != nil {
		panic(err)
	}
	fmt.Println("Results:")
	doc.Find("#modelPartSearchResults .modelSearchResultsItemLeft p span").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").Attr("href")
		if url != "" {
			fmt.Println(i+1, "-->", url)
			fmt.Println(i+1, "-->", s.Find("a").Text())
		}
	})
}
