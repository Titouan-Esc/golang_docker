package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	db := DB()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var user User
		rows, err := db.Query(`SELECT * FROM "users"`)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for rows.Next() {
			if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
				log.Fatalln(err.Error())
			}
		}

		fmt.Fprintf(res, "%v", user)
	})

	if err := http.ListenAndServe(":4200", nil); err != nil {
		log.Fatalln(err.Error())
	}
}

func DB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public", "db", "5432", "admin", "aaa111AAA", "postgres")
	log.Println(dsn)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return conn
}
