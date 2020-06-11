package main

import (
	"net/http"
	"strconv"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Response - Export for testing
type Response struct {
	Error  bool   `json:"error"`
	String string `json:"string"`
	Answer int64  `json:"answer"`
}

//Error - Export for testing
type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// var router *gin.Engine

func main() {
	router := SetupRouter()
	router.Run("0.0.0.0:80")
}

//SetupRouter - Export router for testing
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// router.Use(cors.Default())
	router.GET("/quotient", handleGet)
	return router
}

func badRequestError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Error{Error: true, Message: message})
}

func handleGet(c *gin.Context) {
	x, found := c.GetQuery("x")
	if !found {
		badRequestError(c, "x parameter not provided")
		return
	} else if _, e := strconv.ParseInt(x, 10, 64); e != nil {
		badRequestError(c, "x parameter is not a valid integer")
		return
	}

	y, found := c.GetQuery("y")
	if !found {
		badRequestError(c, "y parameter not provided")
		return
	} else if _, e := strconv.ParseInt(y, 10, 64); e != nil {
		badRequestError(c, "y parameter is not a valid integer")
		return
	} else if y == "0" {
		badRequestError(c, "Divisor (y) must not be 0")
		return
	}

	params := c.Request.URL.Query()
	delete(params, "x")
	delete(params, "y")
	if len(params) > 0 {
		badRequestError(c, "Unrecognised parameter(s) provided")
		return
	}

	xInt, _ := strconv.ParseInt(x, 10, 64)
	yInt, _ := strconv.ParseInt(y, 10, 64)

	answer := xInt / yInt

	res := Response{
		Error:  false,
		String: x + " / " + y + " = " + strconv.FormatInt(answer, 10),
		Answer: answer,
	}

	c.JSON(http.StatusOK, res)
}
