package main

import (
	"log"
	"net/http"
	"twitterClone/users"

	"twitterClone/tweets"
)

func main() {
	log.Printf("Up and Running")
	http.HandleFunc("/users/", users.Users)
	http.HandleFunc("/tweets/", tweets.Tweets)

	http.ListenAndServe(":2828", nil)
}
