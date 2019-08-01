package main

import (
	"fmt"
	"os"
	"z/poster"
)

func main() {
	apikey := os.Getenv("POSTER_API_KEY")

	if apikey == "" {
		fmt.Println("missing POSTER_API_KEY")
		fmt.Println("export POSTER_API_KEY=<YOUR_POSTER_API_KEY>")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Println("missing keyword")
		fmt.Println("usage: go run main.go <KEYWORD>")
		os.Exit(1)
	}

	movie, err := poster.FetchMovie(apikey, os.Args[1])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	err = poster.FetchPosterImage(movie.Poster, movie.Title)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}
