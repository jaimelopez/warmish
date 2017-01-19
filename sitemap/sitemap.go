package sitemap

import (
	"encoding/xml"
	"net/http"
	"reflect"
)

type sitemapElement struct {
	Location     string  `xml:"loc"`
	Priority     float32 `xml:"priority"`
	Modification string  `xml:"lastmod"`
}

type SitemapIndex struct {
	sitemapElement
	Sitemaps []Sitemap `xml:"sitemap"`
}

type Sitemap struct {
	sitemapElement
	Urls []Url `xml:"url"`
}

type Url struct {
	sitemapElement
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

func (sitemap *Sitemap) AddUrl(url Url) {
	sitemap.Urls = append(sitemap.Urls, url)
}

func (sitemap *Sitemap) AddUrlCollection(urls []Url) {
	for _, url := range urls {
		sitemap.AddUrl(url)
	}
}

func Fetch(locations []string) SitemapIndex {
	sitemapIndex := SitemapIndex{}

	for _, location := range locations {
		url := Url{}
		url.Location = location

		sitemap := Sitemap{}
		sitemap.AddUrl(url)

		sitemapIndex.AddSitemap(sitemap)
	}

	Crawl(&sitemapIndex)

	return sitemapIndex
}

func Crawl(sitemapIndex *SitemapIndex) error {
	for _, currentSitemap := range sitemapIndex.Sitemaps {
		sitemapIndex, sitemap, error := Parse(currentSitemap.Location)

		if error != nil {
			return error
		}

		if reflect.DeepEqual(sitemapIndex, SitemapIndex{}) {
			sitemapIndex.AddSitemapCollection(sitemapIndex.Sitemaps)
		} else if reflect.DeepEqual(sitemap, Sitemap{}) {
			currentSitemap.AddUrlCollection(sitemap.Urls)
		}
	}

	return nil
}

func Parse(location string) (sitemapIndex SitemapIndex, sitemap Sitemap, error error) {
	response, error := http.Get(location)

	if error != nil {
		return
	}

	defer response.Body.Close()

	error = xml.NewDecoder(response.Body).Decode(&sitemapIndex)

	error = xml.NewDecoder(response.Body).Decode(&sitemap)

	return
}
