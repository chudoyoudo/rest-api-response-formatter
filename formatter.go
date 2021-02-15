package rest_api_response_formatter

func GetResponseData(data interface{}, errors *map[string][]string) *map[string]interface{} {
    return &map[string]interface{}{
        "data":   data,
        "errors": *errors,
    }
}
