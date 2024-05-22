package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type responseObject struct {
	Id           string   `json:"__id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"author_slug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"date_added"`
	DateModified string   `json:"date_modified"`
}

func main() {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response responseObject
	if err = json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}
	fmt.Print(response.Content)
}
