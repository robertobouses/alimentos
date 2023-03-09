package main

import "net/http"

func createAlimento(c *gin.Context) {
	var alimento Alimento
	err := c.BindJSON(&alimento)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO alimentos(id, nombre, calorias, recomendacion) values(?, ?, ?, ?)")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(alimento.Nombre, alimento.Calorias, alimento.Recomendacion)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Alimento creado exitosamente"})
}
