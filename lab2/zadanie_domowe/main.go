package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	LoadEnvVariables()
}

func currencyValues(r *gin.Engine) {
	r.POST("/submit", func(c *gin.Context) {
		amount := c.PostForm("amount")
		currency1 := c.PostForm("curr1")
		currency2 := c.PostForm("curr2")

		if amount == "" && currency1 == "" && currency2 == "" {
			fmt.Println("Error: amount, currency1 and currency2 is required")
			return
		}
		converter := NewCurrencyConverter(currency1, currency2, amount)

		c.HTML(200, "index.html", gin.H{
			"BestPrice":    fmt.Sprintf("%.8f", converter.bestPrice),
			"WorstPrice":   fmt.Sprintf("%.8f", converter.worstPrice),
			"AveragePrice": fmt.Sprintf("%.8f", converter.averagePrice),
			"Result":       fmt.Sprintf("%.4f", converter.finalValue),
			"Selling":      currency1,
			"Buying":       currency2,
			"Success":      true,
		})
	})
}

func createHtml(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	createHtml(r)
	currencyValues(r)
	r.Run()
}
