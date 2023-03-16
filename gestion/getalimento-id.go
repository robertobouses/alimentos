package gestion

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlimentoID(c *gin.Context) {
	id := c.Param("id")

	db, err := sql.Open("postgres", "postgres://roberto:pass1234@localhost:4444/postgres?sslmode=disable")

	var alimento Alimento
	err = db.QueryRow("SELECT id, nombre, calorias, recomendacion FROM alimentos WHERE rowid=?", id).Scan(&alimento.ID, &alimento.Nombre, &alimento.Calorias, &alimento.Recomendacion)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	alimentoJson, err := json.Marshal(alimento)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json", alimentoJson)
}
