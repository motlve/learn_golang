package main

import (
	"bufio"
	"fmt"
	"modularization-go/01-basics/day-4/modularization-go/grades"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("------- CEK NILAI DISINI -------")
	fmt.Println()

	fmt.Println("================================")
	fmt.Print("Masukan nama siswa: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("================================")

	fmt.Println("Lihat semua nilai? (Iya/Tidak)")
	seeAllStudy, _ := reader.ReadString('\n')
	seeAllStudy = strings.TrimSpace(seeAllStudy)

	fmt.Println("================================")

	if seeAllStudy == "Iya" {
		grades.CheckAllGrades(name)
	} else {
		fmt.Print("Masukan mata pelajaran: ")
		study, _ := reader.ReadString('\n')
		study = strings.TrimSpace(study)
		fmt.Println("================================")

		fmt.Println("========================")
		fmt.Println("Hasil Pengecekan Nilai")
		fmt.Println("========================")

		fmt.Println()

		grades.CheckGrades(name, study)
	}

}
