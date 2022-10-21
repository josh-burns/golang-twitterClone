package tweets

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"
)

func getTweets(id string) string {
	var tweetArray []string
	DbAccessString := GoDotEnvVariable("DB_ACCESS_STRING")
	db, err := sql.Open("mysql", DbAccessString)

	if err != nil {
		log.Fatal(err)
	}

	query := "SELECT * FROM Twitter.tweets WHERE authorId = " + id + ";"

	rows, _ := db.Query(query)
	defer rows.Close()

	log.Printf("getting tweets for userID " + id)

	for rows.Next() {
		SingleTweet := new(Tweet)
		err := rows.Scan(
			&SingleTweet.Id,
			&SingleTweet.AuthorId,
			&SingleTweet.DateTweeted,
			&SingleTweet.TweetBody,
			&SingleTweet.Likes,
			&SingleTweet.Retweets,
		)

		if err != nil {
			log.Fatal(err)
		}

		marshalled, _ := json.Marshal(SingleTweet)

		tweetArray = append(tweetArray, string(marshalled))

	}
	stringArray := strings.Join(tweetArray, ",")
	return string(stringArray)
}
