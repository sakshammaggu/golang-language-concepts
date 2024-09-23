package main

import (
	"fmt"
	"net/url"
)

const urlLink="https://www.example.com/path/to/resource?param1=value1&param2=value2#fragment"

func parseUrl(urlLink string) {
	parsedUrl, err:=url.Parse(urlLink)
	if (err!=nil) {
		panic(err)
	}
	fmt.Printf("Type of parsed url: %T\n", parsedUrl)
	fmt.Println("Protocol:", parsedUrl.Scheme)
	fmt.Println("Hostname:", parsedUrl.Hostname())
	fmt.Println("Query:", parsedUrl.Query())
	fmt.Println("Fragment:", parsedUrl.Fragment)
}

func main() {
	fmt.Println("Handling url in go.")

	// parse url
	parseUrl(urlLink)
}