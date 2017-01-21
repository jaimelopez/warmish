package sitemap

import (
	"errors"
)

type Sitemap struct {
	sitemapElement
	Urls     []Url `xml:"url"`
	Hydrated bool
}

func (this *Sitemap) AddUrl(url Url) *Sitemap {
	this.Urls = append(this.Urls, url)

	return this
}

func (this *Sitemap) AddUrlCollection(urls []Url) *Sitemap {
	for _, url := range urls {
		this.AddUrl(url)
	}

	return this
}

func (this *Sitemap) GetAllLocations() []string {
	locations := []string{}

	for _, url := range this.Urls {
		locations = append(locations, url.Location)
	}

	return locations
}

func (this *Sitemap) Hydrate() (error error) {
	if this.Hydrated {
		return
	}

	if this.Location == "" {
		error = errors.New("Sitemap can't be hydrated: missing location")
		return
	}

	_, parsedSitemap, error := Parse(this.Location)

	if error != nil {
		return
	}

	*this = parsedSitemap
	this.MarkAsHydrated()

	return
}

func (this *Sitemap) MarkAsHydrated() *Sitemap {
	this.Hydrated = true

	return this
}
