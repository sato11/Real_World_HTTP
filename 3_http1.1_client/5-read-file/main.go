package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resp, err := http.Post("http://localhost:18888", "text/html", file)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status:", resp.Status)
}
