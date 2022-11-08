package tweets

import (
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
	}

	os.Setenv("ENV_FILE_LOCATION", "../env/.env")

	// fill mock db with expected row
	tweetRows := sqlmock.NewRows([]string{"id", "authorId", "dateTweeted", "tweetBody", "likes", "retweets"}).
		AddRow(1, 81, "2022-03-27T12:06:50Z", "hello", 3, 9)
	likeRows := sqlmock.NewRows([]string{"id", "tweetId", "likerId", "dateLiked"}).
		AddRow(1, 20, 88, "2022-03-27T12:06:50Z")

	// Stub the insert operation
	mock.ExpectExec("INSERT into Twitter.likes").WillReturnResult(sqlmock.NewResult(1, 1))

	// Stub the checks after insert
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Twitter.tweets WHERE id = 20")).WillReturnRows(tweetRows)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Twitter.likes WHERE tweetId = 20 AND likerId = 88;")).WillReturnRows(likeRows)

	mock.ExpectCommit()

	// mock the req body
	bodyString := "{\"tweetId\": 20, \"likerId\": 20}"

	r := io.NopCloser(strings.NewReader(bodyString))

	res := LikeTweet(db, r)

	expected := "liked"

	if res != expected {
		t.Error("\n EXPECTED: \n", expected, "\n RECIEVED: \n", res)
	}

}
