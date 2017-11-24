package utils

import (
	"errors"
	"fmt"
)

func NormalizeObject(data map[interface{}]interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	for key, value := range data {
		stringKey, success := key.(string)
		if !success {
			return nil, errors.New(fmt.Sprintf("Key was not a string: %v\n", key))
		}

		if mapValue, success := value.(map[interface{}]interface{}); success {
			normalized, err := NormalizeObject(mapValue)
			if err != nil {
				return nil, err
			}
			out[stringKey] = normalized
		} else if arrayValue, success := value.([]interface{}); success {
			normalized, err := NormalizeArray(arrayValue)
			if err != nil {
				return nil, err
			}
			out[stringKey] = normalized
		} else {
			out[stringKey] = value
		}
	}
	return out, nil
}

func NormalizeArray(data []interface{}) ([]interface{}, error) {
	out := make([]interface{}, 0, len(data))
	for _, value := range data {
		if mapValue, success := value.(map[interface{}]interface{}); success {
			normalized, err := NormalizeObject(mapValue)
			if err != nil {
				return nil, err
			}
			out = append(out, normalized)
		} else if arrayValue, success := value.([]interface{}); success {
			normalized, err := NormalizeArray(arrayValue)
			if err != nil {
				return nil, err
			}
			out = append(out, normalized)
		} else {
			out = append(out, value)
		}
	}
	return out, nil
}
