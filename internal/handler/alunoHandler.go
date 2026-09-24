package handler

import (
	"api-gin/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	alunos    []models.Aluno
	proximoID int
)

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	for _, a := range alunos {
		if a.Matricula == aluno.Matricula {
			c.JSON(http.StatusConflict, gin.H{"erro": "matrícula já cadastrada"})
			return
		}
	}

	proximoID++
	aluno.Id = proximoID
	alunos = append(alunos, aluno)

	c.JSON(http.StatusCreated, aluno)
}

func ListarAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, alunos)
}

func BuscarAluno(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	for _, a := range alunos {
		if a.Id == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "aluno não encontrado"})
}

func AtualizarAluno(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	var dados models.Aluno
	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	for i, a := range alunos {
		if a.Id == id {
			dados.Id = id
			alunos[i] = dados
			c.JSON(http.StatusOK, dados)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "aluno não encontrado"})
}

func ExcluirAluno(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	for i, a := range alunos {
		if a.Id == id {
			alunos = append(alunos[:i], alunos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensagem": "aluno excluído"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "aluno não encontrado"})
}
