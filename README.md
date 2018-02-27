# PS Plus Games API

A Go API to get the current month's free PlayStation Plus games

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Install Go following the [Getting Started section of the Go Programming Language site](https://golang.org/doc/install).

### Installing

You can either build it using:

```go
go build
```

Or simply run it with:

```go
go run main.go
```

## Endpoints

There's only one endpoint:

```
localhost:8000/free-games
```

When calling it, it will return a JSON response, with an array containing the games. Each game has the same format:

```
{
	"title": "Title of the game",
	"console": "Console for this game (for example, PS4)"
}
```

### Extra configuration

In OS X there's a message that pops up every time you run a Go app that has outbound connections.
To avoid having that displayed constantly when running it locally, you can set the env var GOENV to 'dev'

```
export GOENV=dev
```

Which will run the server while disabling that warning.

## Deployment

[Use the Official Go Docker image](https://hub.docker.com/_/golang/).

## Uses

* [gorilla/mux](https://github.com/gorilla/mux) - Go router and dispatcher for requests
* [goquery](https://github.com/PuerkitoBio/goquery) - Html parser and CSS selectors, used for scraping

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
