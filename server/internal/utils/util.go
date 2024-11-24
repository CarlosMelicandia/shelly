package utils

import (
	"database/sql"
	"strconv"
)

func ToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{String: "", Valid: false}
}

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
