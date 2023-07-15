# SiteMap Generator

## What is SiteMap?
Sitemaps are an easy way for webmasters to inform search engines about pages on their sites that are available for crawling. 

For more info: https://www.sitemaps.org/index.html

## How to run it locally?

Run
```
go run cmd/*.go
```


## Functionality

- Uses BFS for faster performance.
- Uses [html-link-parser](github.com/tusharr-patil/html-link-parser) package to parse the link from html page.
- Sample output given [above](https://github.com/tusharr-patil/sitemap/blob/main/test.xml).
