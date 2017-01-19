package sitemap

type SitemapIndex struct {
	sitemapElement
	Sitemaps []Sitemap `xml:"sitemap"`
}

func (sitemapIndex *SitemapIndex) AddSitemap(sitemap Sitemap) {
	sitemapIndex.Sitemaps = append(sitemapIndex.Sitemaps, sitemap)
}

func (sitemapIndex *SitemapIndex) AddSitemapCollection(sitemaps []Sitemap) {
	for _, sitemap := range sitemaps {
		sitemapIndex.AddSitemap(sitemap)
	}
}

func (sitemapIndex *SitemapIndex) GetAllUrls() []Url {
	urls := []Url{}

	for _, sitemap := range sitemapIndex.Sitemaps {
		for _, url := range sitemap.Urls {
			urls = append(urls, url)
		}
	}

	return urls
}
