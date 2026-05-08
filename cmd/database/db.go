package database

import (
	"GOeGIN/cmd/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=ROOT password=ROOT dbname=ROOT"
	DB, err := gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
