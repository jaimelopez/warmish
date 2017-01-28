# Warmish
This is a simple project that allows the warm-up of a website thru their sitemaps.

It was thought to works over a Varnish instance in order to having all the website updated for example either after a server restart or websites with a lot of fragmentation in their pages and many dispersion in the visits of those pages.

The idea could be that you have a configured Varnish instance with a certain document life-time period and you want to have truly updated thoses documents. When you visit a out of date document, then varnish request to the server for a new version of these document. If you should like have an updated version all the time, then you could schedule a job launching Warmish which will do a warm-up of your site.  

## Installation
```
> go get github.com/jaimelopez/warmish
> warmish --help
```

## Configuration
You can specifiy a mixing list specifing as well, sitemap indexex (which include other sitemaps) and normal sitemaps.
```
sitemaps:
  - "http://www.your-website.com/sitemap-index-1.xml"
  - "http://www.your-website.com/sitemap-index-2.xml"
  - "http://www.your-website.com/other-sitemap.xml"
```

The number of concurrent requests and the lapse time between those group of requests could be specified like the next example:
```
concurrency: 5
break: 1000ms
```

Warmish can refresh a url (purge + warm-up) or just one of these...
Maybe you only want to warm-up the urls or perhaps you want refresh the urls, invalidating and warming up them. In the last case the applicattion will do twice requests one for each action.

In many cases, Varnish could be configured to refresh a document after the purge request of it (like the Varnish's config file showed below). That's why these options are.
```
purge: true
warmup: true
```

You can run the applicattion just once or configuring it for schedules cron. If you want to schedules the execution then you can configure it.

For example, if you would to execute it every hour:
```
schedule: "0 0 * * * *"
```
Or in the same way but more simple:
```
schedule: "@hourly"
```
You can also specify a periods like:
```
schedule: "@every 4h"
```

## Enabling Varnish purge
If you just want to allows the warm-up and don't want to deal with the purge of those files then you can leave Varnish doing how it was configured.

In our case, we want to allow it and the following example shows how configure it for that.

First we need to specify certain addresses from which the purge will be enabled, and then the section in where we allows to do that purge:
```
acl purgers {
    "localhost";
    "192.168.1.0"/24;
}
 
sub vcl_recv {
    if (req.method == "PURGE") {
        if (!client.ip ~ purgers) {
            return (synth(405, "Purging not allowed for " + client.ip));
        }

        return (purge);
    }
}
```

In the following section we setting upIn the following section we set up that Varnish always will preload a document after the purge of it.
>Notice that with this option you could set to false the ***warmup*** option in Warnish because Varnish itself will warmup the documents after the purge.

```
sub vcl_purge {
    set req.method = "GET";
    return (restart);
}
```