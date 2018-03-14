# Dialogflow Go Webhook

![Go Version](https://img.shields.io/badge/go-1.9-brightgreen.svg)
![Go Version](https://img.shields.io/badge/go-1.10-brightgreen.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/leboncoin/dialogflow-go-webhook)](https://goreportcard.com/report/github.com/leboncoin/dialogflow-go-webhook)
[![Build Status](https://drone.depado.eu/api/badges/leboncoin/dialogflow-go-webhook/status.svg)](https://drone.depado.eu/leboncoin/dialogflow-go-webhook)
[![codecov](https://codecov.io/gh/leboncoin/dialogflow-go-webhook/branch/master/graph/badge.svg)](https://codecov.io/gh/leboncoin/dialogflow-go-webhook)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/leboncoin/dialogflow-go-webhook/blob/master/LICENSE)
[![Godoc](https://godoc.org/github.com/leboncoin/dialogflow-go-webhook?status.svg)](https://godoc.org/github.com/leboncoin/dialogflow-go-webhook)


Simple library to create compatible DialogFlow v2 webhooks using Go.

This package is only intended to create webhooks, it doesn't implement the whole 
DialogFlow API.

<!-- TOC -->

- [Dialogflow Go Webhook](#dialogflow-go-webhook)
- [Introduction](#introduction)
    - [Goal of this package](#goal-of-this-package)
    - [Disclaimer](#disclaimer)
- [Installation](#installation)
    - [Using dep](#using-dep)
    - [Using go get](#using-go-get)
- [Usage](#usage)
    - [Handling incoming request](#handling-incoming-request)
    - [Retrieving params and contexts](#retrieving-params-and-contexts)
    - [Responding with a fulfillment](#responding-with-a-fulfillment)
- [Examples](#examples)

<!-- /TOC -->

# Introduction

## Goal of this package

This package aims to implement a complete way to receive a DialogFlow payload,
parse it, and retrieve data stored in the parameters and contexts by providing
your own data structures. 

It also allows you to format your response properly, the way DialogFlow expects
it, including all the message types and platforms. (Such as cards, carousels,
quick replies, etc…)

## Disclaimer

As DialogFlow's v2 API is still in Beta, there may be breaking changes. If 
something breaks, please file an issue.

# Installation

## Using dep

If you're using [dep](https://github.com/golang/dep) which is the recommended
way to vendor your dependencies in your project, simply run this command :

`dep ensure -add github.com/leboncoin/dialogflow-go-webhook`

## Using go get

If your project isn't using `dep` yet, you can use `go get` to install this
package :

`go get github.com/leboncoin/dialogflow-go-webhook`

# Usage

Import the package `github.com/leboncoin/dialogflow-go-webhook` and use it with
the `dialogflow` package name. To make your code cleaner, import it as `df`.
All the following examples and usages use the `df` notation.

## Handling incoming request

In this section we'll use the [gin](https://github.com/gin-gonic/gin) router as
it has some nice helper functions that will keep the code concise. For an
example using the standard `http` router, see 
[this example](https://github.com/leboncoin/dialogflow-go-webhook/blob/master/examples/http).

When DialogFlow sends a request to your webhook, you can unmarshal the incoming
data to a `df.Request`. This, however, will not unmarshal the contexts and the
parameters because those are completely dependent on your data models.

```go
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

	if err = c.BindJSON(&dfr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func main() {
	r := gin.Default()
	r.POST("/webhook")
	if err := r.Run("127.0.0.1:8001"); err != nil {
		panic(err)
	}
}
```

## Retrieving params and contexts

```go
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
}
```

In this example we're getting the DialogFlow request and unmarshalling the
params to a defined struct. This is why `json.RawMessage` is used in both the 
`Request.QueryResult.Parameters` and in the `Request.QueryResult.Contexts`.

This also allows you to filter and route according to the `action` and `intent`
DialogFlow detected, which means that depending on which action you detected,
you can unmarshal the parameters and contexts to a completely different data
structure.

The same thing can be done for contexts :

```go
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

	if err = dfr.GetContext("my-awesome-context", &p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
```

## Responding with a fulfillment

DialogFlow expects you to respond with what is called a [fulfillment](https://dialogflow.com/docs/reference/api-v2/rest/v2beta1/WebhookResponse).

This package supports every rich response type.

```go
func HandleWebhook(c *gin.Context) {
	var err error
	var dfr *df.Request
	var p params

	if err = c.BindJSON(&dfr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Send back a fulfillment
	dff := &df.Fulfillment{
		FulfillmentMessages: df.Messages{
			df.ForGoogle(df.SingleSimpleResponse("hello", "hello")),
			{RichMessage: df.Text{Text: []string{"hello"}}},
		},
	}
	c.JSON(http.StatusOK, dff)
}
```


# Examples

- [Using Gin](https://github.com/leboncoin/dialogflow-go-webhook/blob/master/examples/gin)
- [Using http](https://github.com/leboncoin/dialogflow-go-webhook/blob/master/examples/http)


