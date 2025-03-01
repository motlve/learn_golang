package models

/**
	Struct dijadikan model di folder models
	Gunakan tag json:"field_name" untuk JSON
*/

type Students struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float64 `json:"score"`
}

