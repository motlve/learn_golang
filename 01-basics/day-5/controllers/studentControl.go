package controllers

import (
	"01-basics/01-basics/day-5/models"
	"fmt"
)

type StudentsController struct {
	// Buar struct baru yang "membungkus" model.Students
	Student models.Students
}

// Tambahkan method ke struct baru ini
func (studentController StudentsController) Display() {
	fmt.Printf("Nama: %s\n, Umur: %d\n, Nilai: %2.f\n", studentController.Student.Name, studentController.Student.Age, studentController.Student.Score)
}
