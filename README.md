# SiteMap Generator

## What is SiteMap?
Sitemaps are an easy way for webmasters to inform search engines about pages on their sites that are available for crawling. 

For more info: https://www.sitemaps.org/index.html

Checkout **Deployed app**: [sitemap generator](https://sitemap-generator.up.railway.app/)

## Functionality

- Uses BFS for faster performance.
- User can add modified date, priority in the xml.
- User can crawl through 500 pages.
- Uses [html-link-parser](github.com/tusharr-patil/html-link-parser) package to parse the link from html page.
- Sample output given [above](https://github.com/tusharr-patil/sitemap/blob/main/test.xml).
  
## How to run it locally?

Run
```
go run cmd/*.go
```
