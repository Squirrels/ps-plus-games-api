package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/PuerkitoBio/goquery"
    "strings"
    "os"
)

// A struct to keep the data for the game (with the custom json render name)
type Game struct {
    Title     string `json:"title"`
    Console   string `json:"console"`
}

func getFreeGames(w http.ResponseWriter, r *http.Request) {
    doc, err := goquery.NewDocument("https://www.playstation.com/en-us/explore/playstation-plus")
    if err != nil {
        log.Fatal(err)
    }
    var free_games []Game
    // Get the element by the contains seach flag, and look for the text
    // Then get the parent of that element's parent
    // Finally, iterate through the li elements in the ul INSIDE the div element
    doc.Find("h3:contains('Membership Includes These Free Games')").First().Parent().Parent().Find("div ul li").Each(func(index int, item *goquery.Selection) {
        // Note that we replace a special character being used in the sony website (U+00A0 : NO-BREAK SPACE [NBSP])
        game_parts := strings.Split(strings.Replace(item.Text(), " ", "", -1), "//")
        title, console := game_parts[0], game_parts[1]
        // Add it to the array while removing trailing whitespace
        free_games = append(free_games, Game{strings.Trim(title, " "), strings.Trim(console, " ")})
    })
    // Time to send it, first marshal the array
    jData, err := json.Marshal(free_games)
    if err != nil {
        panic(err)
        return
    }
    // A correct response
    w.WriteHeader(http.StatusOK)
    // ... which contains a JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/free-games", getFreeGames).Methods("GET")
    // Added a check for the env variable GOENV, if it's "dev", then add localhost to the server listener
    // This is so that the annoying "accept incoming connections?" message is not displayed in OS X
    if os.Getenv("GOENV") == "dev" {
        log.Fatal(http.ListenAndServe("localhost:8000", router))
    } else{
        log.Fatal(http.ListenAndServe(":8000", router))
    }
    
}