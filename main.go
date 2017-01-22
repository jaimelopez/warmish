package main

import (
	"warmish/config"
	"warmish/sitemap"
	"warmish/warmer"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var configFile string

	app := cli.NewApp()
	app.Name = "Warmish"
	app.Usage = "Warm-up a website thru their sitemaps"
	app.Author = "Jaime Lopez"
	app.Version = "0.1"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config",
			Value: "config.yml",
			Usage: "specify the configuration file",
			Destination: &configFile,
		},
	}

	app.Action = func(c *cli.Context) error {
		configuration := config.New(configFile)

		sitemapIndex := sitemap.Crawl(configuration.Sitemaps)
		urls := sitemapIndex.GetAllLocations()

		(warmer.Warmer{
			Purge: configuration.Purge,
			Warmup: configuration.Warmup,
			Concurrency: configuration.Concurrency,
			Break: configuration.Break,
		}).Run(urls)

		return nil
	}

	app.Run(os.Args)
}
