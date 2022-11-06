package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func getUserByEmail(db *sql.DB, email string) string {
	var marshalled []byte
	retrievedUser := new(User)

	query := "SELECT * FROM Twitter.users WHERE email = \"" + email + "\";"

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
