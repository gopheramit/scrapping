package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

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

}
