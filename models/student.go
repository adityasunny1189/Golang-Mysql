package models

import (
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
)

type Student struct {
	Name   string
	Rollno int
}

func (s *Student) AddDetails(db *sql.DB, in *bufio.Scanner) {
	fmt.Printf("Enter name: ")
	in.Scan()
	s.Name = in.Text()
	fmt.Printf("Enter rollno: ")
	in.Scan()
	s.Rollno, _ = strconv.Atoi(in.Text())
	_, err := db.Exec("INSERT INTO students (name, rollno) VALUES (?, ?)", s.Name, s.Rollno)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Student) ShowDetails() {
	fmt.Printf("Name: %s\t Roll: %d\n", s.Name, s.Rollno)
}
