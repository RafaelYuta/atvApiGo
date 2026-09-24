package models

type Alocacao struct {
	TurmaId int    `json:"turma_id"`
	SalaId  int    `json:"sala_id"`
	Dia     string `json:"dia"`
	Inicio  string `json:"inicio"`
	Fim     string `json:"fim"`
}
