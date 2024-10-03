package models

type Question struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Detail      string   `json:"detail"`
	Author      string   `json:"author"`
	CreatedAt   string   `json:"created_at"`
	AuthorEmail string   `json:"author_email"`
	Answers     []Answer `json:"answers"`
}

type Answer struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	AuthorEmail string `json:"author_email"`
	AuthorName  string `json:"author_name"`
	QuestionId  int    `json:"question_id,omitempty"`
	IsBest      bool   `json:"is_best,omitempty"`
}
