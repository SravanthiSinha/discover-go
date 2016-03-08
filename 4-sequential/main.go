package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

type mv struct {
	Title      string  `json:"Title"`
	Year       string  `json:"Year"`
	Rated      string  `json:"Rated"`
	Released   string  `json:"Released"`
	Runtime    string  `json:"Runtime"`
	Genre      string  `json:"Genre"`
	Director   string  `json:"Director"`
	Writer     string  `json:"Writer"`
	Actors     string  `json:"Actors"`
	Plot       string  `json:"Plot"`
	Language   string  `json:"Language"`
	Country    string  `json:"Country"`
	Awards     string  `json:"Awards"`
	Poster     string  `json:"Poster"`
	Metascore  string  `json:"Metascore"`
	ImdbRating float64 `json:"imdbRating,string"`
	ImdbVotes  string  `json:"imdbVotes"`
	ImdbID     string  `json:"imdbID"`
	Type       string  `json:"Type"`
	Response   string  `json:"Response"`
}

type SearchList struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   int    `json:"Year,string"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults int    `json:"totalResults,string"`
	Response     string `json:"Response"`
}

func main() {

	var movie string
	flag.StringVar(&movie, "movie", "Batman", "a string var")
	flag.Parse()
	client := &http.Client{}
	start := time.Now()
	resp, _ := client.Get("http://www.omdbapi.com/?s=" + movie)

	searchlist := new(SearchList)
	json.NewDecoder(resp.Body).Decode(&searchlist)

	for index := 0; index < len(searchlist.Search); index++ {
		client2 := &http.Client{}

		resp2, _ := client2.Get("http://www.omdbapi.com/?i=" + searchlist.Search[index].ImdbID + "&plot=short&r=json")

		m := new(mv)
		json.NewDecoder(resp2.Body).Decode(&m)
		fmt.Println("The movie :", m.Title, " was released in", m.Year, " - the IMBD rating is ", m.ImdbRating*10, "% with ", m.ImdbVotes, "votes")
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Seconds()

	fmt.Printf("execution time is %vs\n", elapsed)

	return

}
