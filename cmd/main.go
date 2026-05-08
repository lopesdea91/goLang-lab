package main

import (
	"GOeGIN/cmd/database"
	"GOeGIN/cmd/models"
	"GOeGIN/cmd/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	models.Alunos = []models.Aluno{
		{Nome: "Anderson", CPF: "074.155.710-02", RG: "15.929.689-4"},
		{Nome: "Maria", CPF: "090.434.780-08", RG: "16.870.336-1"},
		{Nome: "João", CPF: "329.177.950-93", RG: "15.623.599-7"},
	}

	routes.HandlerRequests()

}
