# Golang Webcrawler

This is a webcrawl to get all links and and their counts throughout a domain.  To initiate the crawl
use the following command structure.

`crawler <baseURL> <maxConcurrency> <maxPages>`

The script will then output to the command line the various pages and their count.

To speed up the script increase the maxconConcurrency.  (careful not to overdo it still)

Setting the maxPages will stop the script early if crawling a large website with lots of pages under the same domain i.e., Reddit (probably don't do this anyway)


### Possible Future Enchancements
- Save the report as a CSV rather than print to console
- Count external links, as well as internal links, and add them to the report
    - just don't crawl the external links
- Use a graphics library to create an image showing the links between pages