package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Alimento struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Calorias      int    `json:"calorias"`
	Recomendacion string `json:"recomendacion"`
}

var db *sql.DB

func main() {
	var err error
	db, err := sql.Open("postgres", "postgres://roberto:pass1234@localhost:4444/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()

	router.POST("/alimentos", createAlimento)
	router.GET("/alimentos/:id", getAlimentoID)
	router.GET("/alimentos/:tipo", getAlimentosTipo)
	//router.GET("/alimentos/hipercaloricos", getAlimentosTipo)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
