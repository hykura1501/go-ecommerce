package pkg

func ResponseSuccess(message string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
	}
}

func ResponseError(message string, err error) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"error":   err.Error(),
	}
}

func ResponseSuccessWithData(message string, extraFields map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{
		"message": message,
	}
	for k, v := range extraFields {
		response[k] = v
	}
	return response
}
