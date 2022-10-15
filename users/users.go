package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id            int
	Bio           string
	Username      string
	DateCreated   string
	DisplayPicUrl string
}

func getUser(id string) string {
	var marshalled []byte
	retrievedUser := new(User)
	db, err := sql.Open("mysql", "root:Cypress123!!@tcp(localhost:3306)/Twitter")

	query := "SELECT * FROM Twitter.users WHERE userId = " + id + ";"

	if err != nil {
		log.Fatal("error initialising connection with DB - ", err)
	}

	rows, _ := db.Query(query)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&retrievedUser.Id,
			&retrievedUser.Bio,
			&retrievedUser.Username,
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
		splitUrl := strings.Split(r.URL.String(), "/")
		userId := splitUrl[len(splitUrl)-1]
		fmt.Fprint(w, getUser(userId))
	case "POST":
		fmt.Println("this is a post request")
	}
}
