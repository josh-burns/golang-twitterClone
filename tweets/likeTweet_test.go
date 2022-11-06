package tweets

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestLikeTweet(t *testing.T) {

	// Set up the mock db
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

		os.Setenv("ENV_FILE_LOCATION", "../env/.env")

		// fill mock db with expected row
		rows := sqlmock.NewRows([]string{"id", "authorId", "dateTweeted", "tweetBody", "likes", "retweets"}).
			AddRow(1, 81, "2022-03-27T12:06:50Z", "hello", 3, 9)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Twitter.tweets WHERE id = 20")).WillReturnRows(rows)
		mock.ExpectCommit()

		mock.ExpectExec("UPDATE Twitter.tweets SET likes = 4 WHERE id = 20;").WillReturnResult(sqlmock.NewResult(1, 1))

		// mock the req body
		bodyString := "{\"tweetId\": 20, \"likerId\": 20}"

		r := io.NopCloser(strings.NewReader(bodyString))
		fmt.Println(r)
		res := LikeTweet(db, r)
		fmt.Fprintln(os.Stderr, "hello")

		expected := "test"
		if res == expected {
			t.Error(res)
		}
		fmt.Println(res)

	}
}
