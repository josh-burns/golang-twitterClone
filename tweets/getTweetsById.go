package tweets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func GetTweetbyId(id int) string {
	var tweetArray []string
	DbAccessString := GoDotEnvVariable("DB_ACCESS_STRING")
	db, _ := sql.Open("mysql", DbAccessString)

	log.Printf("getting tweet by id %d ...", id)

	query := fmt.Sprintf("SELECT * FROM Twitter.tweets WHERE Id = %v;", id)

	rows, _ := db.Query(query)
	defer rows.Close()

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
	fmt.Println(stringArray, len(stringArray))
	return string(stringArray)
}
