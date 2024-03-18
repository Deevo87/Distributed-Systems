package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"slices"
)

var apiKeys []string

func generateApiKey() string {
	apiKey := uuid.New()
	return apiKey.String()
}

func sendApiKey(r *gin.Engine) {
	r.GET("/sendApiKey", func(c *gin.Context) {
		apiKey := generateApiKey()
		apiKeys = append(apiKeys, apiKey)
		c.JSON(http.StatusOK, gin.H{"apiKey": apiKey})
	})
}

func ifApiKeyExists(apiKey string) bool {
	return slices.Contains(apiKeys, apiKey)
}

func sendTradeInfo(r *gin.Engine) {
	r.GET("/convert", func(c *gin.Context) {
		amount := c.Query("amount")
		curr1 := c.Query("curr1")
		curr2 := c.Query("curr2")
		apiKey := c.Query("apiKey")

		if !ifApiKeyExists(apiKey) {
			c.JSON(401, gin.H{"errorMessage": "There is no such api key."})
			return
		}

		converter := NewCurrencyConverter(curr1, curr2, amount)

		if converter.bestPrice == 0 && converter.finalValue == 0 {
			c.JSON(400, gin.H{
				"error":   "Bad Request",
				"message": "Invalid input data",
			})
			return
		}

		c.JSON(200, gin.H{
			"BestPrice":    converter.bestPrice,
			"WorstPrice":   converter.worstPrice,
			"AveragePrice": converter.averagePrice,
			"Result":       converter.finalValue,
			"Selling":      curr1,
			"Buying":       curr2,
			"Success":      true,
		})
	})
}

func main() {
	r := gin.Default()
	sendApiKey(r)
	sendTradeInfo(r)
	r.Run(":3000")
}
