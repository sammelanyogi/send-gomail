package entity

type Message struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"email"`
	Text  string `json:"text"`
}
