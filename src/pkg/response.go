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

func ResponseSuccessWithData(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":    data,
	}
}
