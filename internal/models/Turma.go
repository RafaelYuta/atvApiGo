package models

type Turma struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Disciplina string  `json:"disciplina"`
	Docente    string  `json:"docente"`
	Alunos     []Aluno `json:"alunos"`
}
