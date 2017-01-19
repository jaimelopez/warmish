package sitemap

type sitemapElement struct {
	Location     string  `xml:"loc"`
	Priority     float32 `xml:"priority"`
	Modification string  `xml:"lastmod"`
}
