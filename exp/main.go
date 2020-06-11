package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	/* _, err = db.Exec(`INSERT INTO users(name, email) VALUES($1,$2)`, "Jon Calhoun", "jon@calhoun.io")*/
	//if err != nil {
	//panic(err)
	//}

	/* var id int*/
	//row := db.QueryRow(`INSERT INTO users(name, email) VALUES($1,$2) RETURNING id`, "Christopher Kim", "chris@cj3kim.com")
	//err = row.Scan(&id)
	//if err != nil {
	//panic(err)
	//}

	var id int
	var name, email string
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users
		WHERE email = $1
		OR ID > $2`,
		"chris@cj3kim.com", 2)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&id, &name, &email)
		fmt.Println("ID:", id, "Name:", name, "Email:", email)
	}

	db.Close()
}
