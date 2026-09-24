package handler

import (
	"strconv"
	"strings"

	"api-gin/internal/models"
)

func buscarTurma(id int) (models.Turma, bool) {
	for _, t := range turmas {
		if t.Id == id {
			return t, true
		}
	}
	return models.Turma{}, false
}

func buscarAluno(id int) (models.Aluno, bool) {
	for _, a := range alunos {
		if a.Id == id {
			return a, true
		}
	}
	return models.Aluno{}, false
}

func buscarSala(id int) (models.Sala, bool) {
	for _, s := range salas {
		if s.Id == id {
			return s, true
		}
	}
	return models.Sala{}, false
}

func contarMatriculados(turmaId int) int {
	total := 0
	for _, m := range matriculas {
		if m.TurmaId == turmaId {
			total++
		}
	}
	return total
}

func alunosMatriculados(turmaId int) []models.Aluno {
	var resultado []models.Aluno
	for _, m := range matriculas {
		if m.TurmaId == turmaId {
			if aluno, ok := buscarAluno(m.AlunoId); ok {
				resultado = append(resultado, aluno)
			}
		}
	}
	return resultado
}

func alunoMatriculado(alunoId, turmaId int) bool {
	for _, m := range matriculas {
		if m.AlunoId == alunoId && m.TurmaId == turmaId {
			return true
		}
	}
	return false
}

func alocacaoDaTurma(turmaId int) (models.Alocacao, bool) {
	for _, a := range alocacoes {
		if a.TurmaId == turmaId {
			return a, true
		}
	}
	return models.Alocacao{}, false
}

func turmaAlocada(turmaId int) bool {
	_, ok := alocacaoDaTurma(turmaId)
	return ok
}

func temConflitoAgenda(alunoId int, aloc models.Alocacao, turmaNovaId int) bool {
	for _, m := range matriculas {
		if m.AlunoId != alunoId || m.TurmaId == turmaNovaId {
			continue
		}
		if outra, ok := alocacaoDaTurma(m.TurmaId); ok {
			if outra.Dia == aloc.Dia && sobrepoe(aloc.Inicio, aloc.Fim, outra.Inicio, outra.Fim) {
				return true
			}
		}
	}
	return false
}

func sobrepoe(inicioA, fimA, inicioB, fimB string) bool {
	return paraMinutos(inicioA) < paraMinutos(fimB) && paraMinutos(fimA) > paraMinutos(inicioB)
}

func paraMinutos(horario string) int {
	partes := strings.Split(horario, ":")
	if len(partes) != 2 {
		return 0
	}
	h, _ := strconv.Atoi(partes[0])
	m, _ := strconv.Atoi(partes[1])
	return h*60 + m
}
