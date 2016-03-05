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

	m := new(movie)
	json.NewDecoder(resp.Body).Decode(&m)
	irating, _ := strconv.ParseFloat(m.ImdbRating, 16)
	fmt.Println("The movie :", m.Title, " was released in", m.Year, " - the IMBD rating is ", irating, "% with ", m.ImdbVotes, " votes")
	return

}
