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
	"time"
)

// Joke Struct
type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	ID        int    `json:"id"`
}

func main() {
	fmt.Println("Getting your joke...")
	url := "https://official-joke-api.appspot.com/random_joke"
	jokeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}
	fmt.Print("\033[H\033[2J") // clear console
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	// use default chrome user agent for Windows 10.

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

	joke1 := Joke{}
	jsonErr := json.Unmarshal(body, &joke1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("Heres your joke!" + "\n" + joke1.Setup)
	fmt.Println("Enter for punchline...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(joke1.Punchline)
}
