// This is a simple program that I built to call the public
// Jokes API and print out a random joke.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
)

// Joke Types
type joke_types []string

// Joke Struct
type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	ID        int    `json:"id"`
}

func main() {
	fmt.Println("Getting your joke...")
	url := "https://official-joke-api.appspot.com/jokes/random/1"
	jokeTypeURL := "https://official-joke-api.appspot.com/types"
	// use default chrome user agent for Windows 10.
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

	jokeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}
	fmt.Print("\033[H\033[2J") // clear console

	// get joke types
	typereq, err := http.NewRequest("GET", jokeTypeURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	typereq.Header.Set("User-Agent", userAgent)

	typeres, getErr := jokeClient.Do(typereq)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if typeres.Body != nil {
		defer typeres.Body.Close()
	}
	typebody, readErr := io.ReadAll(typeres.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	types := joke_types{}
	jsonErrtyp := json.Unmarshal(typebody, &types)
	if jsonErrtyp != nil {
		log.Fatal(jsonErrtyp)
	}

	fmt.Println("What joke type?" + "\n" + strings.Join(types[:], ","))
	selected_type := bufio.NewScanner(os.Stdin)
	selected_type.Scan()

	selectionExists := (slices.Contains(types, selected_type.Text()))

	// Get Jokes
	// add joke type.
	if !(selected_type.Err() != nil) && (len(selected_type.Text()) > 0) && selectionExists {
		url = "https://official-joke-api.appspot.com/jokes/" + selected_type.Text() + "/random/"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", userAgent)

	res, getErr := jokeClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jokes := []Joke{}
	jsonErr := json.Unmarshal(body, &jokes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if len(jokes) <= 0 {
		fmt.Println("Failed to get your joke :( !!!")
		fmt.Println("API RESPONSE: " + "\n" + string(body))
	} else {
		fmt.Println("Heres your joke!" + "\n" + jokes[0].Setup)
		fmt.Println("Enter for punchline...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		fmt.Println(jokes[0].Punchline)
	}

}
