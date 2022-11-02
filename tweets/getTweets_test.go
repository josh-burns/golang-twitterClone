package tweets

import (
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFindByID(t *testing.T) {
	// Set up the mock db
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	os.Setenv("ENV_FILE_LOCATION", "../env/.env")

	// fill mock db with expected row
	rows := sqlmock.NewRows([]string{"id", "authorId", "dateTweeted", "tweetBody", "likes", "retweets"}).
		AddRow(1, 81, "2022-03-27T12:06:50Z", "hello", 3, 9)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Twitter.tweets WHERE authorId = 81")).WillReturnRows(rows)
	mock.ExpectCommit()

	// call the method
	res := GetTweets(db, "81")

	expected := "{\"Id\":1,\"AuthorId\":81,\"DateTweeted\":\"2022-03-27T12:06:50Z\",\"TweetBody\":\"hello\",\"Likes\":3,\"Retweets\":9}"

	// fail test if method response doesn't match
	if res != expected {
		t.Error("\n EXPECTED: \n", expected, "\n GOT: \n ", res)
	}
}
