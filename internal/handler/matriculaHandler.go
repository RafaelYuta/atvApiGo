package handler

import (
	"net/http"
	"strconv"

	"api-gin/internal/models"

	"github.com/gin-gonic/gin"
)

func MatricularAluno(c *gin.Context) {
	turmaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	if _, ok := buscarTurma(turmaId); !ok {
		c.JSON(http.StatusNotFound, gin.H{"erro": "turma não encontrada"})
		return
	}

	var body struct {
		AlunoId int `json:"aluno_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	if _, ok := buscarAluno(body.AlunoId); !ok {
		c.JSON(http.StatusNotFound, gin.H{"erro": "aluno não encontrado"})
		return
	}

	if alunoMatriculado(body.AlunoId, turmaId) {
		c.JSON(http.StatusConflict, gin.H{"erro": "aluno já matriculado nesta turma"})
		return
	}

	if aloc, ok := alocacaoDaTurma(turmaId); ok {
		sala, _ := buscarSala(aloc.SalaId)
		if contarMatriculados(turmaId) >= sala.Capacidade {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": "capacidade da sala excedida"})
			return
		}
		if temConflitoAgenda(body.AlunoId, aloc, turmaId) {
			c.JSON(http.StatusConflict, gin.H{"erro": "conflito de agenda do aluno"})
			return
		}
	}

	matriculas = append(matriculas, models.Matricula{AlunoId: body.AlunoId, TurmaId: turmaId})
	c.JSON(http.StatusCreated, gin.H{"mensagem": "aluno matriculado"})
}

func ListarAlunosTurma(c *gin.Context) {
	turmaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	if _, ok := buscarTurma(turmaId); !ok {
		c.JSON(http.StatusNotFound, gin.H{"erro": "turma não encontrada"})
		return
	}

	c.JSON(http.StatusOK, alunosMatriculados(turmaId))
}
