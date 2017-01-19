package sitemap

type Sitemap struct {
	sitemapElement
	Urls []Url `xml:"url"`
}

func (sitemap *Sitemap) AddUrl(url Url) {
	sitemap.Urls = append(sitemap.Urls, url)
}

func (sitemap *Sitemap) AddUrlCollection(urls []Url) {
	for _, url := range urls {
		sitemap.AddUrl(url)
	}
}
