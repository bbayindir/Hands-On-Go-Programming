package main

import (_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
)

type Profile struct{
	ProfileId int
	FirstName string
	LastName string
	Age int
}

func main(){
	db, err := sql.Open("sqlite3", "./personal.db")
	checkError(err)

	var profile Profile
	rows, err := db.Query("select ProfileId, FirstName, LastName, Age from Profile where FirstName = ? and LastName = ?","Tarik", "Guney")
	checkError(err)

	for rows.Next(){
		err := rows.Scan(&profile.ProfileId, &profile.FirstName, &profile.LastName, &profile.Age)
		checkError(err)

		fmt.Println(profile)
	}

	rows.Close()
	db.Close()
}

func checkError(err error) {
	if (err != nil) {
		panic(err)
	}
}
