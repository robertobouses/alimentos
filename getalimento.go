package main

import (
	"encoding/json"
	"net/http"
)

func getAlimento(c *gin.Context) {
	id := c.Param("id")

	var alimento Alimento
	err := db.QueryRow("SELECT id, nombre, calorias, recomendacion FROM alimentos WHERE rowid=?", id).Scan(&alimento.ID, &alimento.Nombre, &alimento.Calorias, &alimento.Recomendacion)
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
