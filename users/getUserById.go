package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func getUserById(id string) string {
	var marshalled []byte
	retrievedUser := new(User)
	DbAccessString := GoDotEnvVariable("DB_ACCESS_STRING")
	fmt.Println(DbAccessString)
	db, err := sql.Open("mysql", DbAccessString)

	query := "SELECT * FROM Twitter.users WHERE userId = " + id + ";"

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
