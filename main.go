package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

//Comment struct to hold individual comments
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func wordCounter(s string) map[string]int {
	words := strings.Fields(s)
	dict := make(map[string]int)
	for _, word := range words {
		dict[word]++
	}
	return dict
}

func main() {
	endpoint := "https://jsonplaceholder.typicode.com/comments"

	request := http.Client{Timeout: time.Second * 4}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := request.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	comments := []Comment{}
	jsonErr := json.Unmarshal(body, &comments)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, comment := range comments {
		count := wordCounter(comment.Body)

		words := map[int][]string{}
		var sorted []int
		for k, v := range count {
			words[v] = append(words[v], k)
		}
		for k := range words {
			sorted = append(sorted, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
		for _, k := range sorted {
			var max = 1
			fmt.Printf("Comment Id: %d Name: %s\n", comment.ID, comment.Name)
			for _, s := range words[k] {
				fmt.Printf("\t %s, %d\n", s, k)
				max++
				if max == 4 {
					fmt.Print("\n")
					break
				}
			}
		}
	}

}