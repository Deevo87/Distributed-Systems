package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	//"zadanie_domowe/api"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var apiKey string

//func currencyValues(r *gin.Engine) {
//	r.POST("/submit", func(c *gin.Context) {
//		amount := c.PostForm("amount")
//		currency1 := c.PostForm("curr1")
//		currency2 := c.PostForm("curr2")
//		statusCode := 200
//		success := true
//
//		if amount == "" && currency1 == "" && currency2 == "" {
//			fmt.Println("Error: amount, currency1 and currency2 is required")
//			success = false
//			statusCode = 400
//			return
//		}
//
//		converter := api.NewCurrencyConverter(currency1, currency2, amount)
//
//		if converter.bestPrice == 0 && converter.averagePrice == 0 && converter.finalValue == 0 {
//			fmt.Println("Wrong input!!!")
//			success = false
//			statusCode = 400
//		}
//
//		c.HTML(statusCode, "index.html", gin.H{
//			"BestPrice":    fmt.Sprintf("%.8f", converter.bestPrice),
//			"WorstPrice":   fmt.Sprintf("%.8f", converter.worstPrice),
//			"AveragePrice": fmt.Sprintf("%.8f", converter.averagePrice),
//			"Result":       fmt.Sprintf("%.4f", converter.finalValue),
//			"Selling":      currency1,
//			"Buying":       currency2,
//			"Success":      success,
//		})
//	})
//}

func createHtml(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
}

func createApiReq(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error during creating http request to API: ", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error during http response from API: ", err)
		return nil, err
	}
	return res, err
}

func getApiKey(r *gin.Engine) {
	r.GET("/getApiKey", func(c *gin.Context) {
		url := "http://localhost:3000/sendApiKey"

		res, err := createApiReq(url)
		if err != nil {
			fmt.Println("Error during creating api request in FxRates API: ", err)
		}

		defer res.Body.Close()
		fmt.Println(res.Body)
		body, _ := ioutil.ReadAll(res.Body)
		var responseData struct {
			Api string `json:"apiKey"`
		}
		fmt.Println(string(body))

		err = json.Unmarshal(body, &responseData)
		if err != nil {
			fmt.Println("Error during decoding JSON, ", err)
		}
		apiKey = responseData.Api
		fmt.Println(apiKey)
	})
}

func tradeCurrency(r *gin.Engine) {
	r.GET("/convert", func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		url := "http://localhost:3000" + c.Request.URL.String() + "&apiKey=" + apiKey

		res, err := createApiReq(url)
		if err != nil {
			fmt.Println("Error during creating api request in FxRates API: ", err)
		}

		defer res.Body.Close()

		if res.Status == "401 Unauthorized" {
			var error401 struct {
				ErrorMsg string `json:"errorMessage"`
			}
			body, _ := ioutil.ReadAll(res.Body)
			err = json.Unmarshal(body, &error401)
			c.JSON(401, gin.H{"error": error401.ErrorMsg})

			return
		} else if res.Status == "400 Bad Request" {
			var error400 struct {
				Error   string `json:"error"`
				Message string `json:"message"`
			}
			body, _ := ioutil.ReadAll(res.Body)
			err = json.Unmarshal(body, &error400)

			c.JSON(401, gin.H{
				"error":   error400.Error,
				"message": error400.Message,
			})

			return
		}

		body, _ := ioutil.ReadAll(res.Body)

		var responseData struct {
			BestPrice    float64 `json:"BestPrice"`
			WorstPrice   float64 `json:"WorstPrice"`
			AveragePrice float64 `json:"AveragePrice"`
			Result       float64 `json:"Result"`
			Selling      string  `json:"Selling"`
			Buying       string  `json:"Buying"`
			Success      bool    `json:"Success"`
		}

		err = json.Unmarshal(body, &responseData)
		if err != nil {
			fmt.Println("Error during decoding JSON, ", err)
		}
		c.HTML(200, "index.html", gin.H{
			"BestPrice":    fmt.Sprintf("%.8f", responseData.BestPrice),
			"WorstPrice":   fmt.Sprintf("%.8f", responseData.WorstPrice),
			"AveragePrice": fmt.Sprintf("%.8f", responseData.AveragePrice),
			"Result":       fmt.Sprintf("%.4f", responseData.Result),
			"Selling":      responseData.Selling,
			"Buying":       responseData.Buying,
			"Success":      responseData.Success,
		})
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	createHtml(r)
	//currencyValues(r)
	getApiKey(r)
	tradeCurrency(r)
	r.Run(":4000")
}
