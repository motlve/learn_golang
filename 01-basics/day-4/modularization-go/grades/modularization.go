package grades

import (
	"fmt"
)

/*
*
Jika kode terlalu panjang, sebaiknya dipisah dalam package terpisah agar lebih rapi.

1️⃣ Buat folder baru, misalnya modularization-go/grades
2️⃣ Di dalamnya, buat file modularization.go
3️⃣ Tambahkan kode berikut di modularization.go
*/

// membuat data nilai dari tiap pelajaran dan tiap siswa menggunakan map[]
var studentData = map[string]map[string]float64{
	"Rizky Aditiyo": {
		"Matematika":       85,
		"Bahasa Indonesia": 90,
		"Bahasa Inggris":   80,
		"IPA":              80,
	},
	"Naifah Salsabila T Ariani": {
		"Matematika":       75,
		"Bahasa Indonesia": 100,
		"Bahasa Inggris":   85,
		"IPA":              80,
	},
	"Jefri Dwi Hartanto": {
		"Matematika":       80,
		"Bahasa Indonesia": 80,
		"Bahasa Inggris":   80,
		"IPA":              80,
	},
}

func CheckGrades(name string, study string) {
	// cek nama apakah ada di map/data siswa nya
	if valueStudy, exist := studentData[name]; exist {
		// cek apakah mata pelajaran ada di data
		if value, exist := valueStudy[study]; exist {
			fmt.Printf("Nilai %s dalam mata pelajaran %s adalah: %2.f\n", name, study, value)
		} else {
			fmt.Printf("Mata Pelajaran %s tidak ditemukan %s.\n", study, name)
		}
	} else {
		fmt.Printf("Nama %s tidak ditemukan.\n", name)
	}
}

func CheckAllGrades(name string) {
	if valueStudy, exist := studentData[name]; exist {
		fmt.Printf("Nilai semua mata pelajaran untuk %s\n", name)
		for study, value := range valueStudy {
			fmt.Printf("- %s: %2.f\n", study, value)
		}
	} else {
		fmt.Printf("Nama %s tidak ditemukan.\n", name)
	}
}
