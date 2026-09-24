package handler

import (
	"net/http"
	"strconv"

	"api-gin/internal/models"

	"github.com/gin-gonic/gin"
)

var (
	turmas         []models.Turma
	proximoIDTurma int
	matriculas     []models.Matricula
	alocacoes      []models.Alocacao
)

func CriarTurma(c *gin.Context) {
	var turma models.Turma
	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	proximoIDTurma++
	turma.Id = proximoIDTurma
	turmas = append(turmas, turma)

	c.JSON(http.StatusCreated, turmaParaJSON(turma))
}

func ListarTurmas(c *gin.Context) {
	resultado := make([]gin.H, 0, len(turmas))
	for _, t := range turmas {
		resultado = append(resultado, turmaParaJSON(t))
	}
	c.JSON(http.StatusOK, resultado)
}

func BuscarTurma(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	turma, ok := buscarTurma(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"erro": "turma não encontrada"})
		return
	}

	c.JSON(http.StatusOK, turmaParaJSON(turma))
}

func AtualizarTurma(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	var dados models.Turma
	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	for i, t := range turmas {
		if t.Id == id {
			dados.Id = id
			turmas[i] = dados
			c.JSON(http.StatusOK, turmaParaJSON(dados))
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "turma não encontrada"})
}

func ExcluirTurma(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	for i, t := range turmas {
		if t.Id == id {
			turmas = append(turmas[:i], turmas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensagem": "turma excluída"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "turma não encontrada"})
}

func turmaParaJSON(t models.Turma) gin.H {
	return gin.H{
		"id":              t.Id,
		"nome":            t.Nome,
		"disciplina":      t.Disciplina,
		"docente":         t.Docente,
		"qntMatriculados": contarMatriculados(t.Id),
		"alocada":         turmaAlocada(t.Id),
	}
}
