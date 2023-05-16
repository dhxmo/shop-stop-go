package utils

import "github.com/gin-gonic/gin"

func Response(data interface{}, message string, err string) gin.H {
	res := gin.H{
		"data": data,
		"msg":  message,
	}

	if err != "" {
		res["error"] = err
	}

	return res
}
