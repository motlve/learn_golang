package controllers

import (
	"01-basics/01-basics/day-5/challenge/test-001/models"
	"fmt"
)

type StudentsController struct {
	StudentController models.Students
}

func (sc StudentsController) Display() {
	fmt.Printf("Name: %s\n, Age: %d\n, Score: %2.f\n", sc.StudentController.Name, sc.StudentController.Age, sc.StudentController.Score)
}
