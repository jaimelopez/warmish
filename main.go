package main

import (
	"errors"
	"os"
	"warmish/config"
	"warmish/sitemap"
	"warmish/warmer"
	"github.com/robfig/cron"
	"github.com/urfave/cli"
)

func main() {
	var configFile string

	app := cli.NewApp()
	app.Name = "Warmish"
	app.Usage = "Warm-up a website thru their sitemaps"
	app.Author = "Jaime Lopez"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        	"config",
			Usage:       	"specify the configuration file",
			Value:	 	"config.yml",
			Destination:	&configFile,
		},
	}

	app.Action = func(context *cli.Context) error {
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			return errors.New("Configuration file not specified or not found!")
		}

		run(config.New(configFile))

		return nil
	}

	app.Run(os.Args)
}

func run(configuration *config.Config) {
	if configuration.Schedule == "" {
		execute(configuration)
		return
	}

	scheduler := cron.New()
	scheduler.AddFunc(configuration.Schedule, func() {
		execute(configuration)
	})
	scheduler.Start()
}

func execute(configuration *config.Config) {
	sitemapIndex := sitemap.Crawl(configuration.Sitemaps)
	urls := sitemapIndex.GetAllLocations()

	(warmer.Warmer{
		Purge:       configuration.Purge,
		Warmup:      configuration.Warmup,
		Concurrency: configuration.Concurrency,
		Break:       configuration.Break,
	}).Run(urls)
}
