package services

import (
	"bufio"
	"database/sql"
	"demo-files/models"
	"fmt"
)

func GetStudents(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var studentlist []models.Student

	for rows.Next() {
		var stu models.Student
		if rerr := rows.Scan(&stu.Name, &stu.Rollno); rerr != nil {
			fmt.Println(rerr)
			return
		}
		studentlist = append(studentlist, stu)
	}
	for _, s := range studentlist {
		s.ShowDetails()
	}
}

func AddStudent(db *sql.DB, in *bufio.Scanner) {
	u := new(models.Student)
	u.AddDetails(db, in)
}
