package tweets

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Tweet struct {
	Id          int
	AuthorId    int
	DateTweeted string
	TweetBody   string
	Likes       int
	Retweets    int
}

type NewTweet struct {
	AuthorId    string
	DateTweeted string
	TweetBody   string
	Likes       int
	Retweets    int
}

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func likeTweet() {

}

func retweetTweet() {

}

func Tweets(w http.ResponseWriter, r *http.Request) {
	splitUrl := strings.Split(r.URL.String(), "/")
	fmt.Println("this is the tweets working")
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, getTweets(splitUrl[len(splitUrl)-1]))
	case "POST":
		if splitUrl[len(splitUrl)-1] == "new" {
			fmt.Println("new tweet incoming")
			fmt.Fprintf(w, newTweet(r.Body))
		}
	}
}
