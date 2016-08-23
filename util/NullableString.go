package util

import (
	"database/sql/driver"
	"errors"
)

type NullableString struct{
  Str *string
}

// Scan - Implement the database/sql scanner interface
func (nstr *NullableString) Scan(value interface{}) error {
	// if value is nil
	if value == nil {
		*((*nstr).Str) = ""
		return nil
	}

  sv, err := driver.String.ConvertValue(value);
  if err == nil {
		// if this is a bool type
    s, ok := sv.(string);
		if  ok {
			// set the value of the pointer yne to YesNoEnum(v)
			*((*nstr).Str) = s
			return nil
		}
	}
	// otherwise, return an error
	return errors.New("failed to scan NullableString")
}
