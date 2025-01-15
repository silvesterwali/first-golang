package utils

import "reflect"

// RemoveField removes specified fields from a struct or map and returns a new map without those fields.
// data can be a struct or map, and fieldToRemove is a list of field names to remove.
func RemoveField(data interface{}, fieldToRemove []string) interface{} {
	val := reflect.ValueOf(data)

	// If the data is a pointer, get the element it points to.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		// Create a result map to store the fields that are not removed.
		result := make(map[string]interface{})

		// Iterate through the fields of the struct.
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldName := field.Name
			fieldValue := val.Field(i).Interface()

			shouldRemove := false

			// Check if the current field is in the list of fields to remove.
			for _, fieldToRemove := range fieldToRemove {
				if fieldName == fieldToRemove {
					shouldRemove = true
					break
				}
			}

			// If the field should not be removed, add it to the result map.
			if !shouldRemove {
				result[fieldName] = fieldValue
			}
		}
		return result

	case reflect.Map:
		// Create a result map to store the key-value pairs that are not removed.
		result := make(map[string]interface{})

		// Iterate through the keys of the map.
		for _, key := range val.MapKeys() {
			keyString := key.String()
			value := val.MapIndex(key).Interface()

			shouldRemove := false

			// Check if the current key is in the list of keys to remove.
			for _, fieldToRemove := range fieldToRemove {
				if keyString == fieldToRemove {
					shouldRemove = true
					break
				}
			}

			// If the key should not be removed, add it to the result map.
			if !shouldRemove {
				result[keyString] = value
			}
		}
		return result

	default:
		// If data is neither a struct nor a map, return it as is.
		return data
	}
}
