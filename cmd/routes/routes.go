package routes

import (
	"GOeGIN/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()

	r.GET("alunos", controllers.ExibeTodosAlunos)

	r.Run(":5000")
}
