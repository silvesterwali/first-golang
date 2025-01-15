package utils

import "github.com/gin-gonic/gin"

func ResponseData(data interface{}) gin.H {
	return gin.H{
		"status": "success",
		"data":   data,
	}
}

func ResponseDataPagination(data interface{}, meta interface{}) gin.H {
	return gin.H{
		"status": "success",
		"data":   data,
		"meta":   meta,
	}
}
