package gestion

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetAlimentosTipo(c *gin.Context) {
	tipo := c.Param("tipo")
	var alimentos []Alimento
	var err error
	db, err := sqlx.Open("postgres", "postgres://roberto:pass1234@localhost:4444/postgres?sslmode=disable")

	switch tipo {
	case "hipocaloricos":
		err = db.Select(&alimentos, "SELECT * FROM alimentos WHERE calorias < 100")
	case "hipercaloricos":
		err = db.Select(&alimentos, "SELECT * FROM alimentos WHERE calorias > 500")
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de alimentos no válido"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alimentos)
}
