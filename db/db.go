package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// database handle
var db *sql.DB

func ConnectDb() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "gakkai",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Database Connected!")
}

func InsertParticipant(mail, password string) error {
	_, err := db.Exec("INSERT INTO participant (mail, password) VALUES(?, ?)", mail, password)
	if err != nil {
		return fmt.Errorf("InsertParticipant: %v", err)
	}
	return nil
}

func IsUniqueParticipant(mail string) (bool, error) {
	var id int
	isUnique := false

	row := db.QueryRow("SELECT id FROM participant WHERE mail = ?", mail)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			isUnique = true
			return isUnique, nil
		}
		return isUnique, fmt.Errorf("IsUniqueParticipant: %v", err)
	}
	return isUnique, nil
}
