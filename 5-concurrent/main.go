package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type mv struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Response   string `json:"Response"`
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

func getMovie(ID string) {
	client := &http.Client{}

	resp, _ := client.Get("http://www.omdbapi.com/?i=" + ID + "&plot=short&r=json")

	defer resp.Body.Close()
	m := new(mv)
	json.NewDecoder(resp.Body).Decode(&m)
	rating, _ := strconv.ParseFloat(m.ImdbRating, 64)
	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes.\n", m.Title, m.Year, int(rating*10), m.ImdbVotes)
	return
}
func main() {

	var movie string
	flag.StringVar(&movie, "movie", "Batman", "a string var")
	flag.Parse()
	client := &http.Client{}
	start := time.Now()
	resp, _ := client.Get("http://www.omdbapi.com/?s=" + movie)

	defer resp.Body.Close()

	searchlist := new(SearchList)

	json.NewDecoder(resp.Body).Decode(&searchlist)

	var wg sync.WaitGroup
	for index := 0; index < len(searchlist.Search); index++ {
		wg.Add(1)
		x := searchlist.Search[index].ImdbID
		go func(x string) {
			defer wg.Done()
			getMovie(x)

		}(x)
	}
	wg.Wait()

	fmt.Printf("execution time is %v\n", time.Since(start))

	return

}
