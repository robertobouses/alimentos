package gestion

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Alimento struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Calorias      int    `json:"calorias"`
	Recomendacion string `json:"recomendacion"`
}

func CalcularRecomendacion(calorias int) string {
	if calorias > 200 {
		return "Baja recomendación en su dieta"
	} else {
		return "Alta recomendación en su dieta"
	}
}

func CreateAlimento(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://roberto:pass1234@localhost:4444/postgres?sslmode=disable")

	var alimento Alimento
	err = c.BindJSON(&alimento)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calcular la recomendación
	alimento.Recomendacion = CalcularRecomendacion(alimento.Calorias)

	result, err := db.Exec("INSERT INTO alimentos (nombre, calorias, recomendacion) VALUES (?, ?, ?)", alimento.Nombre, alimento.Calorias, alimento.Recomendacion)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()

	alimento.ID = int(id)

	c.JSON(http.StatusOK, alimento)
}
