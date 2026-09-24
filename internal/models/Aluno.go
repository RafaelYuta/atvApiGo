package models

type Aluno struct {
	Id        int    `json:"id"`
	Matricula int    `json:"matricula"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
}
