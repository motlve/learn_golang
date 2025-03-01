package main

import (
	"01-basics/01-basics/day-5/models"
	"01-basics/01-basics/day-5/utils"
	"fmt"
)

/*
* Entry Point

	* Penjelasan Output
	1️⃣ JSON: {"name":"Rizky Aditiyo","age":22,"score":88.5}
💡 Apa yang terjadi?

Struct Student diconvert ke format JSON pakai json.Marshal().
Hasilnya berupa string JSON.
	2️⃣ Data dari file: {"name":"Rizky Aditiyo","age":22,"score":88.5}
💡 Apa yang terjadi?

JSON disimpan ke file student.json pakai os.WriteFile()
JSON dibaca dari file pakai os.ReadFile(), hasilnya masih sama.
	3️⃣ Decoded: {Rizky Aditiyo 22 88.5}
💡 Apa yang terjadi?

JSON dibaca dari file lalu diubah kembali ke struct pakai json.Unmarshal().
Hasilnya berupa struct Student.		

📌 Kesimpulan
Encode JSON ➝ Struct → JSON string
Simpan & Baca JSON dari file ➝ JSON disimpan & dibaca ulang
Decode JSON ➝ JSON string → struct Go
*/

func main() {
	student := models.Students{Name: "Rizky Aditiyo", Age: 22, Score: 88.5}

	// encode JSON
	jsonData, _ := utils.EncodeJSON(student)
	fmt.Println("JSON: ", jsonData)

	// simpan ke file
	utils.SaveToFile("student.json", jsonData)

	// baca dari file
	data, _ := utils.ReadFromFile("student.json")
	fmt.Println("Data dari file: ", data)

	// decode JSON
	decodedStudent, _ := utils.DecodeJSON(data)
	fmt.Println("Decoded: ", decodedStudent)
}
