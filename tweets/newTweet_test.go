package tweets

import (
	"github.com/DATA-DOG/go-sqlmock"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestNewTweet(t *testing.T) {

	// Set up the mock db

	db, mock, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	os.Setenv("ENV_FILE_LOCATION", "../env/.env")

	mock.ExpectBegin()
	tweetRows := sqlmock.NewRows([]string{"id", "authorId", "dateTweeted", "tweetBody", "likes", "retweets"}).
		AddRow(30, 81, "2022-03-27T12:06:50Z", "this is a mocked tweet row", 3, 9)

	mock.ExpectExec("INSERT into Twitter.tweets").WillReturnResult(sqlmock.NewResult(30, 1))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Twitter.tweets")).WillReturnRows(tweetRows)

	mock.ExpectCommit()

	reqBody := "{\"authorId\" : 12,\"tweetedAt\": \"2022-03-27T12:06:50Z\",\"tweetBody\": \"hello\"}"

	r := io.NopCloser(strings.NewReader(reqBody))

	res := NewTweet(db, r)

	expected := "test"

	if res != expected {
		t.Error("\n EXPECTED: \n", expected, "\n RECIEVED: \n", res)
	}
}
