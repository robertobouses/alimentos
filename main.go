package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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
	db, err = sql.Open("sqlite3", "./alimentos.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	fmt.Println(router)
	router.POST("/alimentos", createAlimento)
	router.GET("/alimentos/:id", getAlimento)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
