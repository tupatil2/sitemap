# SiteMap Generator

## What is SiteMap?
Sitemaps are an easy way for webmasters to inform search engines about pages on their sites that are available for crawling. 

https://www.sitemaps.org/index.html

## How to run it locally?

Import the package:

```
go get -u github.com/tusharr-patil/sitemap
```
Run 

```
Sitemap.GenerateSiteMap(<any site-name in string>)
```

## Functionality

- Using BFS for faster performance
- Using github.com/tusharr-patil/html-link-parser to parse the link from html page
