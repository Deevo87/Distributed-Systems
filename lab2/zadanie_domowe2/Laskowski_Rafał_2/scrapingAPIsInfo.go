package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"sync"
)

type CurrencyConverter struct {
	convertFrom  string
	convertTo    string
	amount       string
	toProcess    []float64
	finalValue   float64
	worstPrice   float64
	bestPrice    float64
	averagePrice float64
	changeChan   chan func(*CurrencyConverter)
	wg           sync.WaitGroup
	doneChan     chan struct{}
}

func NewCurrencyConverter(convertFrom, convertTo, amount string) *CurrencyConverter {
	c := &CurrencyConverter{
		convertFrom: convertFrom,
		convertTo:   convertTo,
		amount:      amount,
		toProcess:   make([]float64, 0),
		changeChan:  make(chan func(converter *CurrencyConverter)),
		doneChan:    make(chan struct{}),
		worstPrice:  math.Inf(1),
		bestPrice:   math.Inf(-1),
	}
	c.wg.Add(2)
	go c.fxRatesApiConv()
	go c.frankfurterApiConv()
	go c.applyChanges()

	c.wg.Wait()
	close(c.doneChan)

	return c
}

func (c *CurrencyConverter) applyChanges() {
	for {
		select {
		case change := <-c.changeChan:
			change(c)
			c.wg.Done()
		case <-c.doneChan:
			return
		}
	}
}

func (c *CurrencyConverter) updateValues(price, result float64) {
	c.changeChan <- func(c *CurrencyConverter) {
		if c.bestPrice < price {
			c.bestPrice = price
			c.finalValue = result
		}
		if c.worstPrice > price {
			c.worstPrice = price
		}
		c.toProcess = append(c.toProcess, price)
		c.averagePrice = (c.averagePrice*(float64(len(c.toProcess))-1) + price) / float64(len(c.toProcess))
	}
}

func (c *CurrencyConverter) Clear() {
	c.toProcess = make([]float64, 0)
}

func (c *CurrencyConverter) createApiReq(url string) (*http.Response, error) {
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

func (c *CurrencyConverter) fxRatesApiConv() (float64, float64, error) {
	url := "https://api.fxratesapi.com/convert?from=" + c.convertFrom + "&to=" + c.convertTo + "&amount=" + c.amount + "&format=json"

	res, err := c.createApiReq(url)
	if err != nil {
		fmt.Println("Error during creating api request in FxRates API: ", err)
		return 0, 0, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var responseData struct {
		Result float64 `json: "result"`
		Info   struct {
			Rate float64 `json: "rate"`
		}
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error during decoding JSON, ", err)
		return -1, -1, err
	}
	fmt.Println(responseData.Result, responseData.Info)

	c.updateValues(responseData.Info.Rate, responseData.Result)

	return responseData.Result, responseData.Info.Rate, nil
}

func (c *CurrencyConverter) frankfurterApiConv() (float64, float64, error) {
	url := "https://api.frankfurter.app/latest?amount=" + c.amount + "&from=" + c.convertFrom + "&to=" + c.convertTo
	res, err := c.createApiReq(url)
	if err != nil {
		fmt.Println("Error during creating api request in FrankFurter API: ", err)
		return 0, 0, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	var responseData struct {
		Rates map[string]float64
	}

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error during decoding JSON in FrankFurter API, ", err)
		return -1, -1, err
	}
	amount, _ := strconv.ParseFloat(c.amount, 64)
	price := responseData.Rates[c.convertTo] / amount
	fmt.Println(responseData.Rates[c.convertTo], price)

	c.updateValues(price, responseData.Rates[c.convertTo])

	return responseData.Rates[c.convertTo], price, nil
}

func (c *CurrencyConverter) LatestApi1(curr1, curr2 string) map[string]float64 {
	url := "https://api.fxratesapi.com/latest?base=" + curr1 + "&currencies=" + curr2 + "&resolution=1m&amount=1&places=6&format=json"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var responseData struct {
		Rates map[string]float64 `json: "rates"`
	}

	err := json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error during decoding JSON, ", err)
		return nil
	}
	fmt.Println(responseData.Rates["USD"])
	return responseData.Rates
}
