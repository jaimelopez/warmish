package main

import (
	"fmt"
	"warmish/sitemap"
	"warmish/config"
)

func main() {
	configuration := config.New("config.yml")

	sitemapIndex := sitemap.Fetch(configuration.Sitemaps)

	fmt.Println(sitemapIndex.GetAllUrls())

	//for _, url := range sitemapIndex.GetAllUrls() {
	//	go callUrl(url)
	//}
}