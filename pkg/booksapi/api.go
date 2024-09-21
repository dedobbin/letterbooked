package booksapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Helper function to pull from a URL and deserialize into JSON, your mileage may vary
func pullFromUrl[T any](url string, v *T) (T, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("HTTP Request Error:", err)
		return *v, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("HTTP Request Error:", err)
		return *v, err
	}

	log.Println("Finished Retrieving data.")
	err = json.Unmarshal(data, v)
	return *v, err
}

// Basic stupid query to get a title, no other params set
func QueryTitle(title string) (result BookQueryResult, err error) {
	log.Println("Sending query for title:", title, "to Archive open Library API")
	params := url.Values{}
	params.Add("q", url.QueryEscape(title))

	queryUrl := "https://openlibrary.org/search.json?" + params.Encode()
	log.Println("Query string generated:", queryUrl)
	return pullFromUrl(queryUrl, &result)
}

func QueryAuthor(author string) (result AuthorQueryResult, err error) {
	log.Println("Sending query for author: ", author, "to Archive open Library API")
	params := url.Values{}
	params.Add("q", url.QueryEscape(author))

	queryUrl := "https://openlibrary.org/search/authors.json?" + params.Encode()
	log.Println("Query string generated:", queryUrl)

	return pullFromUrl(queryUrl, &result)
}
