package routes

import (
	"GOeGIN/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()

	r.GET("alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run(":5000")
}
