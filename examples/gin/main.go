package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	df "github.com/leboncoin/dialogflow-go-webhook"
)

type params struct {
	City   string `json:"city"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func HandleWebhook(c *gin.Context) {
	var err error
	var dfr *df.Request
	var p params

	if err = c.BindJSON(&dfr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Retrieve the params of the request
	if err = dfr.GetParams(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Retrieve a specific context
	if err = dfr.GetContext("my-awesome-context", &p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Do things with the context you just retrieved

	// Send back a fulfillment
	dff := &df.Fulfillment{
		FulfillmentMessages: df.Messages{
			df.ForGoogle(df.SingleSimpleResponse("hello", "hello")),
			{RichMessage: df.Text{Text: []string{"hello"}}},
		},
	}
	c.JSON(http.StatusOK, dff)
}

func main() {
	r := gin.Default()
	r.POST("/webhook")
	if err := r.Run("127.0.0.1:8001"); err != nil {
		panic(err)
	}
}
