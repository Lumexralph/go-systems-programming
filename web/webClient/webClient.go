package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	arguments := os.Args

	if len(arguments) != 2 {
		fmt.Printf("Usage: %s URL \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	URL, err := url.Parse(arguments[1])
	if err != nil {
		fmt.Printf("Error parsing the raw URL %s", err)
		os.Exit(1)
	}

	c := &http.Client{}

	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		fmt.Println("An error occurred ", err.Error())
		os.Exit(1)
	}

	httpResponse, err := c.Do(request)
	if err != nil {
		fmt.Println("An error occurred from the response ", err.Error(), URL)
		os.Exit(100)
	}

	fmt.Println("Status Code: ", httpResponse.Status)
	header, _ := httputil.DumpResponse(httpResponse, false)
	fmt.Print(string(header))

	contentType := httpResponse.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	fmt.Println("Character Set: ", characterSet[1])

	if httpResponse.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength is : ", httpResponse.ContentLength)
	}

	length := 0
	var buffer [1024]byte

	r := httpResponse.Body
	defer r.Close()

	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}

	fmt.Println("Response data length: ", length)
}
