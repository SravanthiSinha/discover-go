package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	client := &http.Client{}

	resp, _ := client.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")

	defer resp.Body.Close()
	m := new(movie)
	json.NewDecoder(resp.Body).Decode(&m)
	rating, _ := strconv.ParseFloat(m.ImdbRating, 64)
	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes\n", m.Title, m.Year, int(rating*10), m.ImdbVotes)
	return

}
