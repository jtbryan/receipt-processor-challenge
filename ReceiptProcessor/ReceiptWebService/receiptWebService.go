package main

import (
	"net/http"
	points "receiptProcessor/Points"
	sql "receiptProcessor/SQL"
	validation "receiptProcessor/Validation"

	"github.com/gin-gonic/gin"
)

func processReceipts(c *gin.Context) {
	receipt, err := validation.ValidateReceipt(c)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	} else {
		uuid := sql.InsertReceipt(points.CalculatePoints(receipt))

		m := make(map[string]string)
		m["id"] = uuid

		c.IndentedJSON(http.StatusOK, m)
	}
}

func getPoints(c *gin.Context) {
	points, err := sql.GetReceipts(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		m := make(map[string]int)
		m["points"] = points

		c.IndentedJSON(http.StatusOK, m)
	}
}

func startSQL() {
	sql.Connect()
}

func startRouter() {
	router := gin.Default()
	router.POST("/receipts/process", processReceipts)
	router.GET("/receipts/:id/points", getPoints)

	router.Run("localhost:9090")
}

func main() {
	startSQL()
	startRouter()
}
