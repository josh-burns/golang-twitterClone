package tweets

import (
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
		AddRow(20, 81, "2022-03-27T12:06:50Z", "hello", 3, 9)
	newLikeRow := sqlmock.NewRows([]string{"id", "tweetId", "likerId", "dateLiked"})

	// Stub the checks after insert
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Twitter.tweets")).WillReturnRows(tweetRows)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Twitter.likes")).WillReturnRows(newLikeRow)

	mock.ExpectExec("INSERT into Twitter.likes").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT into Twitter.tweets(authorId, dateTweeted, tweetBody, likes, retweets)").WithArgs("[]").WillReturnResult(sqlmock.NewResult(1, 1))

	// mock the req body
	bodyString := "{\"tweetId\": 20, \"likerId\": 20}"

	r := io.NopCloser(strings.NewReader(bodyString))

	res := LikeTweet(db, r)

	expected := "liked"
	
	if res != expected {
		t.Error("\n EXPECTED: \n", expected, "\n RECIEVED: \n", res)
	}

}
