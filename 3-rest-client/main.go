package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	resp, _ := client.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")

	m := new(movie)
	json.NewDecoder(resp.Body).Decode(&m)
	fmt.Println("The movie :", m.Title, " was released in", m.Year, " - the IMBD rating is ", m.ImdbRating*10, "% with ", m.ImdbVotes, "votes")
	return

}
