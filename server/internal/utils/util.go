package utils

import (
	"strconv"
)

func ConvertToIntField(rawBody *map[string]interface{}, fieldName string) error {
	value, ok := (*rawBody)[fieldName].(string)
	if ok {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		(*rawBody)[fieldName] = intValue
	}
	return nil
}
