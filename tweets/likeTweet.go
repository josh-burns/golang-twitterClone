package tweets

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

func isTweetAlreadyLiked(tweetId int, likerId int) bool {

	var hasAlreadyLiked bool

	DbAccessString := GoDotEnvVariable("DB_ACCESS_STRING")
	db, _ := sql.Open("mysql", DbAccessString)

	log.Printf("Checking if user has already liked tweet ...")

	query := fmt.Sprintf("SELECT * FROM Twitter.likes WHERE tweetId = %v AND likerId = %v;", tweetId, likerId)

	rows, _ := db.Query(query)
	defer rows.Close()

	count := 0
	for rows.Next() {
		count += 1
	}

	if count > 0 {
		hasAlreadyLiked = true
		log.Printf("User has already Liked tweet ... ")
	} else {
		hasAlreadyLiked = false
	}
	return hasAlreadyLiked
}

func likeTweet(body io.ReadCloser) string {
	bytes, _ := io.ReadAll(body)
	jsonString := string(bytes)

	likerId := gjson.Get(jsonString, "likerId")
	tweetId := gjson.Get(jsonString, "tweetId")
	tweetIdNum, _ := strconv.Atoi(tweetId.Raw)

	if isTweetAlreadyLiked(tweetIdNum, int(likerId.Num)) == false {
		tweetToLike := GetTweetbyId(tweetIdNum)

		currentNumberOfLikes := gjson.Get(tweetToLike, "Likes")
		likesNum, _ := strconv.Atoi(currentNumberOfLikes.Raw)

		likesNum = likesNum + 1

		DbAccessString := GoDotEnvVariable("DB_ACCESS_STRING")
		db, _ := sql.Open("mysql", DbAccessString)

		dateLiked := time.Now().Format(time.RFC3339)

		tweetTableLikeQuery := fmt.Sprintf("UPDATE Twitter.tweets SET likes = %d WHERE id = %d;", likesNum, tweetIdNum)
		likesTableLikeQuery := fmt.Sprintf("INSERT into Twitter.likes (tweetId, likerId, dateLiked) values ('%d', '%v', '%s')", tweetIdNum, likerId, dateLiked)

		tweetTableRes, err := db.Exec(tweetTableLikeQuery)
		likesTableRes, err := db.Exec(likesTableLikeQuery)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tweetTableRes)
		fmt.Println(likesTableRes)

		lastId, _ := likesTableRes.LastInsertId()
		log.Printf("Tweet liked : ID = %d, by userID %v", lastId, likerId)

		return "liked"

	} else {
		return "duplicateTweet"
	}

	return "like function complete"
}
