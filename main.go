package main

import (
	"fmt"
	"net/http"
	"twitterClone/users"
)

func main() {
	fmt.Println("working")

	http.HandleFunc("/users/", users.Users)

	http.ListenAndServe(":2828", nil)
}
