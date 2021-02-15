package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User model
type User struct {
	ID              uint64
	Name            string
	Email           string
	password        string
	EmailVerifiedAt string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// FindUser gets a user by Identifier
func FindUser(id string) *User {
	db, dberr := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/creator_core")
	if dberr != nil {
		log.Fatal(dberr)
	}
	defer db.Close()
	stmt, err := db.Prepare("select `id`, `name`, `email`, `password` from `users` where `name` = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		u := new(User)

		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.password); err != nil {
			log.Fatal(err)
		}
		return u
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// Constructor for a new User
func newUser(n string, e string, p string) *User {
	u := new(User)
	u.Name = n
	u.Email = e
	u.password = p
	u.CreatedAt = time.Now()
	u.EmailVerifiedAt = ""
	u.UpdatedAt = time.Now()
	return u
}
