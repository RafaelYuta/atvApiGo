package handler

import (
	"net/http"
	"strconv"

	"api-gin/internal/models"

	"github.com/gin-gonic/gin"
)

func AlocarSalaTurma(c *gin.Context) {
	turmaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	if _, ok := buscarTurma(turmaId); !ok {
		c.JSON(http.StatusNotFound, gin.H{"erro": "turma não encontrada"})
		return
	}

	var body models.Alocacao
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	sala, ok := buscarSala(body.SalaId)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"erro": "sala não encontrada"})
		return
	}

	if contarMatriculados(turmaId) > sala.Capacidade {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": "capacidade da sala menor que o número de alunos"})
		return
	}

	for _, a := range alocacoes {
		if a.SalaId == body.SalaId && a.Dia == body.Dia &&
			sobrepoe(body.Inicio, body.Fim, a.Inicio, a.Fim) {
			c.JSON(http.StatusConflict, gin.H{"erro": "sobreposição de horário na sala"})
			return
		}
	}

	for _, aluno := range alunosMatriculados(turmaId) {
		if temConflitoAgenda(aluno.Id, body, turmaId) {
			c.JSON(http.StatusConflict, gin.H{"erro": "conflito de agenda do aluno"})
			return
		}
	}

	body.TurmaId = turmaId
	alocacoes = append(alocacoes, body)
	c.JSON(http.StatusOK, body)
}
