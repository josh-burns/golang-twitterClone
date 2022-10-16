package users

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

func newUser(body io.ReadCloser) string {
	fmt.Println("heyhey")
	fmt.Println("heyhey")
	bytes, _ := io.ReadAll(body)

	jsonString := string(bytes)

	username := gjson.Get(jsonString, "username")
	email := gjson.Get(jsonString, "email")
	displayPicUrl := gjson.Get(jsonString, "displayPicUrl")

	userToCreate := User{
		Email:         email.Raw,
		Username:      username.Raw,
		DateCreated:   time.Now().Format(time.RFC3339),
		DisplayPicUrl: displayPicUrl.Raw,
	}

	db, err := sql.Open("mysql", "root:Cypress123!!@tcp(localhost:3306)/Twitter")

	query := "insert into users (email, username, dateCreated, displayPic) values ('" + userToCreate.Email + "', '" + userToCreate.Username + "', '" + userToCreate.DateCreated + "', '" + userToCreate.DisplayPicUrl + " ');"

	if err != nil {
		log.Fatal("error initialising connection with DB - ", err)
	}

	res, err := db.Exec(query)

	if err != nil {
		log.Default()
	}
	lastId, _ := res.LastInsertId()
	log.Printf("User added : ID = %d", lastId)

	newUserCheckIfAddedResult := getUserById(strconv.Itoa(int(lastId)))

	return newUserCheckIfAddedResult
}
