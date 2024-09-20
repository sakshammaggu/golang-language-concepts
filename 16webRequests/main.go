package main

import (
	"fmt"
	"io"
	"net/http"
)

const url="http://127.0.0.1:5500/16webRequests/index.html"

func main() {
	fmt.Println("Handling web requests in go.")

	// making get web request to url
	response, err1:=http.Get(url)
	if (err1 != nil) { 
		panic(err1)
	} else { 
		fmt.Println("response: ", response) 
		defer response.Body.Close()
	}

	// getting response body
	dataBytes, err2:=io.ReadAll(response.Body)
	if (err2 != nil) { 
		panic(err2)
	} else { 
		data:=string(dataBytes)
		fmt.Println("Data:", data)
	}
}