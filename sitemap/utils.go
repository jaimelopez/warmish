package sitemap

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func Crawl(locations []string) SitemapIndex {
	sitemapIndex := SitemapIndex{}

	for _, location := range locations {
		sitemapIndex.Compose(location)
	}

	return sitemapIndex
}

func Parse(location string) (sitemapIndex SitemapIndex, sitemap Sitemap, error error) {
	response, error := http.Get(location)

	if error != nil {
		return
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	error = xml.Unmarshal(body, &sitemapIndex)
	error = xml.Unmarshal(body, &sitemap)

	return
}
