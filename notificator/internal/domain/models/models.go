package models

type Message struct {
	From  string `json:"from"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Text  string `json:"text"`
}

type RequestBody struct {
	ChatID    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
	Text      string `json:"text"`
}
