package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"twitterClone/users"

	"twitterClone/tweets"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	envFileLocation := os.Getenv("ENV_FILE_LOCATION")
	fmt.Println(envFileLocation)

	if envFileLocation == "" {
		envFileLocation = "env/.env"
	}

	err := godotenv.Load(envFileLocation)

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

func main() {

	DbAccessString := GoDotEnvVariable("DB_ACCESS_STRING")
	db, err := sql.Open("mysql", DbAccessString)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	log.Printf("Up and Running")
	// TODO - refactor users like tweets on 47
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) { users.Users(db, w, r) })
	http.HandleFunc("/tweets/", func(w http.ResponseWriter, r *http.Request) { tweets.Tweets(db, w, r) })

	http.ListenAndServe(":2828", nil)
}
