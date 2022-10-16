package main

import (
	"log"
	"net/http"
	"twitterClone/users"
)

func main() {
	log.Printf("Up and Running")
	http.HandleFunc("/users/", users.Users)

	http.ListenAndServe(":2828", nil)
}
