package database

import (
	"log"

	//Implementado por Almada
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	//Alterado por Almada
	//stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	stringDeConexao := "host="+os.Getenv("HOST")+
	" user="+os.Getenv("USER")+
	" password="+os.Getenv("PASSWORD")+
	" dbname="+os.Getenv("DBNAME")+
	" port="+os.Getenv("PORT")+
	" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
