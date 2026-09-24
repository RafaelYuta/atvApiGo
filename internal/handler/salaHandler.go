package handler

import (
	"api-gin/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	salas         []models.Sala
	proximoIDSala int
)

func CriarSala(c *gin.Context) {
	var sala models.Sala
	if err := c.ShouldBindJSON(&sala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	if sala.Capacidade <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "capacidade deve ser maior que zero"})
		return
	}

	proximoIDSala++
	sala.Id = proximoIDSala
	salas = append(salas, sala)

	c.JSON(http.StatusCreated, sala)
}

func ListarSalas(c *gin.Context) {
	c.JSON(http.StatusOK, salas)
}

func BuscarSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	for _, s := range salas {
		if s.Id == id {
			c.JSON(http.StatusOK, s)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "sala não encontrada"})
}

func AtualizarSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	var dados models.Sala
	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
		return
	}

	for i, s := range salas {
		if s.Id == id {
			dados.Id = id
			salas[i] = dados
			c.JSON(http.StatusOK, dados)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "sala não encontrada"})
}

func ExcluirSala(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}

	for i, s := range salas {
		if s.Id == id {
			salas = append(salas[:i], salas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensagem": "sala excluída"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": "sala não encontrada"})
}
