package tweets

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func GetTweetsTest(t *Testing.T) {
	fmt.Println("oh hey")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "authorId", "dateTweeted", "tweetBody", "likes", "retweets"}).
		AddRow(1, 11, "2021-12-28T22:54:20Z", "hello", 0, 1).
		AddRow(2, 11, "2021-12-28T22:55:20Z", "world", 1, 0)

	w := httptest.NewRecordor()
	mock.ExpectExec("SELECT * FROM Twitter.tweets WHERE authorId = 11;").WillReturnRows(sqlmock.NewRows())

	if err = GetTweets(11); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
