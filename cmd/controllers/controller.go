package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"nome": "anderson",
	})
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(http.StatusOK, gin.H{
		"API diz:": "Olá " + nome + " !! "})
}
