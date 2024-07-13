package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.DefaultClient.Get("https://github.com/")
	if err != nil{
		log.Fatal(err)
	}

	defer resp.Body.Close()


	fmt.Println("Response status:", resp.StatusCode)


	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)

	}

	fmt.Println(string(body))
}

