package tweets

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tidwall/gjson"
)

 
func NewTweet(db *sql.DB, body io.ReadCloser) string {
	bytes, _ := io.ReadAll(body)
	jsonString := string(bytes)

	authorId := gjson.Get(jsonString, "authorId")
	tweetBody := gjson.Get(jsonString, "tweetBody")

	newTweet := NewTweetRequest{
		AuthorId:    authorId.Raw,
		DateTweeted: time.Now().Format(time.RFC3339),
		TweetBody:   tweetBody.Raw,
		Likes:       0,
		Retweets:    0,
	}

	// DbAccessString := GoDotEnvVariable("DB_ACCESS_STRING")
	// db, err := sql.Open("mysql", DbAccessString)

	query := fmt.Sprintf("INSERT into Twitter.tweets (authorId, dateTweeted, tweetBody, likes, retweets) values (%v, '%s', '%s', %v, %v);",
		newTweet.AuthorId,
		newTweet.DateTweeted,
		newTweet.TweetBody,
		newTweet.Likes,
		newTweet.Retweets)

	res, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	lastId, _ := res.LastInsertId()
	log.Printf("Tweet added : ID = %d", lastId)

	newTweetCheckIfAddedResult := GetTweetbyId(db, int(lastId))

	return newTweetCheckIfAddedResult

}
