package sitemap

import (
	"encoding/xml"
	"net/http"
	"reflect"
)

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
