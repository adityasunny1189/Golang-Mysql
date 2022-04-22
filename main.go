package main

import (
	"bufio"
	"database/sql"
	"demo-files/services"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "human",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	in := bufio.NewScanner(os.Stdin)
loop:
	for {
		fmt.Printf("1: Add Student Data\n2. Get Students Data\n3. Exit\nenter your choice: ")
		in.Scan()
		ch := in.Text()
		switch ch {
		case "1":
			services.AddStudent(db, in)
		case "2":
			services.GetStudents(db)
		case "3":
			break loop
		}
	}
}
