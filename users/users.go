package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

type User struct {
	Id            int
	Email         string
	Username      string
	DateCreated   string
	DisplayPicUrl string
}

func getUserById(id string) string {
	var marshalled []byte
	retrievedUser := new(User)
	db, err := sql.Open("mysql", "root:Cypress123!!@tcp(localhost:3306)/Twitter")

	query := "SELECT * FROM Twitter.users WHERE userId = " + id + ";"

	if err != nil {
		log.Fatal("error initialising connection with DB - ", err)
	}

	rows, _ := db.Query(query)
	defer rows.Close()
	fmt.Println(rows)
	for rows.Next() {
		err := rows.Scan(
			&retrievedUser.Id,
			&retrievedUser.Username,
			&retrievedUser.Email,
			&retrievedUser.DateCreated,
			&retrievedUser.DisplayPicUrl,
		)

		if err != nil {
			log.Fatal(err)
		}

		marshalled, err = json.Marshal(retrievedUser)
		if err != nil {
			fmt.Println(err)
		}

	}
	return string(marshalled)
}

func getUserByEmail(email string) string {
	var marshalled []byte
	retrievedUser := new(User)
	db, err := sql.Open("mysql", "root:Cypress123!!@tcp(localhost:3306)/Twitter")

	query := "SELECT * FROM Twitter.users WHERE email = \"" + email + "\";"

	if err != nil {
		log.Fatal("error initialising connection with DB - ", err)
	}

	rows, _ := db.Query(query)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&retrievedUser.Id,
			&retrievedUser.Username,
			&retrievedUser.Email,
			&retrievedUser.DateCreated,
			&retrievedUser.DisplayPicUrl,
		)

		if err != nil {
			log.Fatal(err)
		}

		marshalled, err = json.Marshal(retrievedUser)
		if err != nil {
			fmt.Println(err)
		}
	}
	return string(marshalled)

}

func newUser() {

}
func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		splitUrl := strings.Split(r.URL.String(), "/")
		userId := splitUrl[len(splitUrl)-1]
		intId, _ := strconv.Atoi(userId)

		if intId > 0 {
			fmt.Fprint(w, getUserById(userId))
		} else {
			if isEmailValid(userId) {
				fmt.Fprint(w, getUserByEmail(userId))
			} else {
				fmt.Fprint(w, "invalid user supplied")
			}
		}

	case "POST":
		fmt.Println("this is a post request")
	}
}
