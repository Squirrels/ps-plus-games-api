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
    var base_url string = "https://store.playstation.com"

    doc, err := goquery.NewDocument(base_url + "/en-us/home/games/psplus")
    if err != nil {
        log.Fatal(err)
    }
    // Get the text 'Free Games' link
    link, err3 := doc.Find("a:contains('Free Games')").First().Attr("href")
    if !err3{
        log.Fatal(err3)
    }
    // Now visit that link
    doc2, err2 := goquery.NewDocument(base_url + link)
    if err2 != nil {
        log.Fatal(err2)
    }

    var free_games []Game
    // Get the base container and the divs inside
    doc2.Find(".grid-cell-container div.ember-view .grid-cell-row__container").Each(func(index int, item *goquery.Selection) {
        // For each div inside
        item.Find(".grid-cell__body").Each(func(index2 int, item2 *goquery.Selection) {
            title   := item2.Find(".grid-cell__title").Text()
            console := item2.Find(".grid-cell__left-detail.grid-cell__left-detail--detail-1").Text()
            free_games = append(free_games, Game{strings.Trim(title, " "), strings.Trim(console, " ")})
        })
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