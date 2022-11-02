package tweets

import (
	"database/sql"
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

func retweetTweet(tweetId int) {

}

func Tweets(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	splitUrl := strings.Split(r.URL.String(), "/")
	w.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, GetTweets(db, splitUrl[len(splitUrl)-1]))
	case "POST":
		if splitUrl[len(splitUrl)-1] == "new" {
			fmt.Println("new tweet incoming")
			fmt.Fprintf(w, newTweet(r.Body))
		}
		if splitUrl[len(splitUrl)-1] == "like" {
			fmt.Println("Liking tweet...")
			likeTweetResponse := likeTweet(r.Body)

			switch likeTweetResponse {
			case "liked":
				fmt.Fprintf(w, "success")

			case "duplicateTweet":
				fmt.Fprintf(w, "Tweet Already liked")

			}
		}
	}
}
