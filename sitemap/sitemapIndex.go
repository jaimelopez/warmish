package sitemap

import (
	"reflect"
)

type SitemapIndex struct {
	sitemapElement
	Sitemaps []Sitemap `xml:"sitemap"`
}

func (this *SitemapIndex) AddSitemap(sitemap Sitemap) *SitemapIndex {
	this.Sitemaps = append(this.Sitemaps, sitemap)

	return this
}

func (this *SitemapIndex) AddSitemapCollection(sitemaps []Sitemap) *SitemapIndex {
	for _, sitemap := range sitemaps {
		this.AddSitemap(sitemap)
	}

	return this
}

func (this *SitemapIndex) GetAllUrls() []Url {
	urls := []Url{}

	for _, sitemap := range this.Sitemaps {
		urls = append(urls, sitemap.Urls...)
	}

	return urls
}

func (this *SitemapIndex) GetAllLocations() []string {
	locations := []string{}

	for _, sitemap := range this.Sitemaps {
		locations = append(locations, sitemap.GetAllLocations()...)
	}

	return locations
}

func (this *SitemapIndex) Compose(url string) (sitemapIndex *SitemapIndex, error error) {
	sitemapIndex = this

	parsedSitemapIndex, parsedSitemap, error := Parse(url)

	if error != nil {
		return
	}

	if !reflect.DeepEqual(parsedSitemapIndex, SitemapIndex{}) {
		this.AddSitemapCollection(parsedSitemapIndex.Sitemaps)
	} else if !reflect.DeepEqual(parsedSitemap, Sitemap{}) {
		parsedSitemap.MarkAsHydrated()
		this.AddSitemap(parsedSitemap)
	}

	for index := range this.Sitemaps {
		sitemap := &this.Sitemaps[index]

		error = sitemap.Hydrate()

		if error != nil {
			continue
		}
	}

	return
}
