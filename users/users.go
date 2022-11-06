package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

// func GoDotEnvVariable(key string) string {
// 	// load .env file
// 	err := godotenv.Load(".env")
//
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
//
// 	return os.Getenv(key)
// }

type NewUser struct {
	Email         string
	Username      string
	DateCreated   string
	DisplayPicUrl string
}

type User struct {
	Id            int
	Email         string
	Username      string
	DateCreated   string
	DisplayPicUrl string
}

func Users(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	splitUrl := strings.Split(r.URL.String(), "/")
	switch r.Method {

	case "GET":
		w.Header().Add("Content-Type", "application/json")
		userId := splitUrl[len(splitUrl)-1]
		intId, _ := strconv.Atoi(userId)

		if intId > 0 {
			fmt.Fprint(w, GetUserById(db, userId))
		} else {
			if isEmailValid(userId) {
				w.WriteHeader(200)
				fmt.Fprint(w, getUserByEmail(db, userId))

			} else {
				fmt.Fprint(w, "invalid user supplied")
			}
		}

	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, newUser(db, r.Body))
	}
}
