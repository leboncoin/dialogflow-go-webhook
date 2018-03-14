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

## Introduction

### Goal of this package

This package aims to implement a complete way to receive a DialogFlow payload,
parse it, and retrieve data stored in the parameters and contexts by providing
your own data structures. 

It also allows you to format your response properly, the way DialogFlow expects
it, including all the message types and platforms. (Such as cards, carousels,
quick replies, etc…)

### Disclaimer

As DialogFlow's v2 API is still in Beta, there may be breaking changes. If 
something breaks, please file an issue.

## Installation

### Using dep

If you're using [dep](https://github.com/golang/dep) which is the recommended
way to vendor your dependencies in your project, simply run this command :

`dep ensure -add github.com/leboncoin/dialogflow-go-webhook`

### Using go get

If your project isn't using `dep` yet, you can use `go get` to install this
package :

`go get github.com/leboncoin/dialogflow-go-webhook`

## Usage

Import the package `github.com/leboncoin/dialogflow-go-webhook` and use it with
the `dialogflow` package name. To make your code cleaner, import it as `df`.
All the following examples and usages use the `df` notation.

### Handling incoming request

When DialogFlow sends a request to your webhook, you can unmarshal the incoming
data to a `df.Request`. 

### Responding with a fulfillment

## Examples

- [Using Gin](https://github.com/leboncoin/dialogflow-go-webhook/blob/master/examples/gin)
- [Using http](https://github.com/leboncoin/dialogflow-go-webhook/blob/master/examples/http)


