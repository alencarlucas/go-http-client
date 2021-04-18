package main

import (
	"fmt"

	"github.com/alencarlucas/go-http-client/gohttp"
)

func main() {
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}
