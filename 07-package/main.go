package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(getPage("https://google.com"))
}
func getPage(url string) string {
	// Send HTTP GET request
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer res.Body.Close()

	// Read the response body
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(content)
}
