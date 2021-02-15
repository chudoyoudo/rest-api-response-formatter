package rest_api_response_formatter

import "github.com/gin-gonic/gin"

func GetResponseData(data interface{}, errors *map[string][]string) *gin.H {
    return &gin.H{
        "data":   data,
        "errors": *errors,
    }
}
