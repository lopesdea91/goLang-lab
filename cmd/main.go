package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExiteTodosAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"nome": "anderson",
	})
}

func main() {
	r := gin.Default()

	r.GET("alunos", ExiteTodosAlunos)

	r.Run(":5000")
}
