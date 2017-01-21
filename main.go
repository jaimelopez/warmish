package main

import (
	"warmish/config"
	"warmish/sitemap"
	"warmish/warmer"
)

func main() {
	configuration := config.New("config.yml")

	sitemapIndex := sitemap.Crawl(configuration.Sitemaps)
	urls := sitemapIndex.GetAllLocations()

	(warmer.Warmer {
		Purge: configuration.Purge,
		Concurrency: configuration.Concurrency,
		Break: configuration.Break,
	}).Run(urls)
}
