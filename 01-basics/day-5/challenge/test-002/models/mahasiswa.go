package models

type DataMahasiswa struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

type ListMahasiswa struct {
	Mahasiswa []DataMahasiswa `json:"mahasiswa"`
}
