package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
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

func getMovie(ID string) {
	client := &http.Client{}

	resp, err := client.Get("http://www.omdbapi.com/?i=" + ID + "&plot=short&r=json")
	if err != nil {
		fmt.Errorf("error parsing body %v", err)
		return
	}
	defer resp.Body.Close()
	m := new(mv)
	json.NewDecoder(resp.Body).Decode(&m)
	fmt.Println("The movie :", m.Title, " was released in", m.Year, " - the IMBD rating is ", m.ImdbRating*10, "% with ", m.ImdbVotes, "votes.")
	return
}
func main() {

	var movie string
	flag.StringVar(&movie, "movie", "Batman", "a string var")
	flag.Parse()
	client := &http.Client{}
	start := time.Now()
	resp, err := client.Get("http://www.omdbapi.com/?s=" + movie)
	if err != nil {
		fmt.Errorf("error parsing body %v", err)
		return
	}
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
