package models

type Message struct {
	System string `json:"system"`
	Phone  string `json:"phone"`
	Square string `json:"square"`
}

type RequestBody struct {
	ChatID    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
	Text      string `json:"text"`
}
