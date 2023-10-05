package config

import (
	api_body "go-demo/pkg/body"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TraceRequest(c *gin.Context) {
	// beforeRequest
	c.Next()
	afterRequest(c)
}

func afterRequest(c *gin.Context) {
	if errRes, exists := c.Get("error"); exists {
		successResponse := api_body.ErrorResponse{
			Status:       "error",
			ErrorMessage: errRes,
		}
		c.JSON(http.StatusOK, successResponse)
	} else if resp, exists := c.Get("response"); exists {
		successResponse := api_body.SuccessResponse{
			Status: "success",
			Data:   resp,
		}
		c.JSON(http.StatusOK, successResponse)
	}
	c.Next()
}
