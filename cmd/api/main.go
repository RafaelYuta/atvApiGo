package main

import (
	"net/http"
	"time"

	"api-gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	// Uso dos Middlewares globais nativos e personalizados
	r.Use(gin.Recovery())

	// 4. Mapeamento de Rotas sob Grupo Versionado
	v1 := r.Group("/api/v1")
	{
		// Monitoramento da API
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":    "healthy",
				"timestamp": time.Now(),
				"version":   "1.0.0",
			})
		})

		// Domínio de Alunos
		v1.POST("/alunos", handler.CriarAluno)
		v1.GET("/alunos", handler.ListarAlunos)
		v1.GET("/alunos/:id", handler.BuscarAluno)
		v1.PUT("/alunos/:id", handler.AtualizarAluno)
		v1.DELETE("/alunos/:id", handler.ExcluirAluno)

		// Domínio de Salas
		v1.POST("/salas", handler.CriarSala)
		v1.GET("/salas", handler.ListarSalas)
		v1.GET("/salas/:id", handler.BuscarSala)
		v1.PUT("/salas/:id", handler.AtualizarSala)
		v1.DELETE("/salas/:id", handler.ExcluirSala)

		// Domínio de Turmas
		v1.POST("/turmas", handler.CriarTurma)
		v1.GET("/turmas", handler.ListarTurmas)
		v1.GET("/turmas/:id", handler.BuscarTurma)
		v1.PUT("/turmas/:id", handler.AtualizarTurma)
		v1.DELETE("/turmas/:id", handler.ExcluirTurma)
		v1.POST("/turmas/:id/matricular", handler.MatricularAluno)
		v1.GET("/turmas/:id/alunos", handler.ListarAlunosTurma)
		v1.POST("/turmas/:id/alocar", handler.AlocarSalaTurma)
	}

	r.Run(":8080")
}
