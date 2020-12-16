package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func postScrape() {
	doc, err := goquery.NewDocument("http://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#main article .entry-title").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})

}

func linkScrape() {
	doc, err := goquery.NewDocument("http://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
	})
}

func main() {
	postScrape()
	linkScrape()
}

/*
func main() {
	//downloadUrl()
	//addr := flag.String("addr", "www.google.com", "Website to be scarepped")
	//flag.Parse()
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequest("GET", "www.google.com", nil)

	if err != nil {
		log.Fatal(err)

	}
	request.Header.Set("User-Agent", "test")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)

	}
	defer response.Body.Close()

	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data", n)
}
func downloadUrl() {
	response, err := http.Get("https://www.devdungeon.com/archive")
	if err != nil {

		log.Fatal(err)
	}
	defer response.Body.Close()

	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}

}*/
