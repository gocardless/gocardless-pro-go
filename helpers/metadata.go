package gocardless

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ToMetadataValue converts any value to a string suitable for metadata.
//
// The GoCardless API requires metadata values to be strings. This function
// handles the conversion from common Go types to strings.
//
// Examples:
//
//	ToMetadataValue(12345)           // "12345"
//	ToMetadataValue(true)            // "true"
//	ToMetadataValue([]string{"a"})   // `["a"]`
//	ToMetadataValue(map[string]string{"key": "val"}) // `{"key":"val"}`
func ToMetadataValue(value interface{}) string {
	if value == nil {
		return "null"
	}

	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		// For complex types (slices, maps, structs), serialize to JSON
		bytes, err := json.Marshal(v)
		if err != nil {
			// Fallback to fmt.Sprintf if JSON marshaling fails
			return fmt.Sprintf("%v", v)
		}
		return string(bytes)
	}
}

// ToMetadata converts a map with mixed value types to a metadata-compatible format.
//
// This is useful for converting a map[string]interface{} to map[string]string
// where all values are automatically converted to strings.
//
// Example:
//
//	metadata := ToMetadata(map[string]interface{}{
//	    "user_id": 12345,
//	    "is_active": true,
//	    "tags": []string{"vip", "premium"},
//	})
//	// Result: map[string]string{
//	//   "user_id": "12345",
//	//   "is_active": "true",
//	//   "tags": `["vip","premium"]`,
//	// }
func ToMetadata(obj map[string]interface{}) map[string]string {
	result := make(map[string]string, len(obj))
	for key, value := range obj {
		result[key] = ToMetadataValue(value)
	}
	return result
}

// IsValidMetadata checks if a map contains only string values.
//
// This can be used to verify that metadata is in the correct format
// before sending it to the API.
func IsValidMetadata(obj map[string]interface{}) bool {
	for _, value := range obj {
		if _, ok := value.(string); !ok {
			return false
		}
	}
	return true
}

// ParseMetadataValue attempts to parse a metadata string value back to a specific type.
//
// Examples:
//
//	var userID int
//	ParseMetadataValue("12345", &userID)  // userID = 12345
//
//	var isActive bool
//	ParseMetadataValue("true", &isActive)  // isActive = true
//
//	var tags []string
//	ParseMetadataValue(`["vip","premium"]`, &tags)  // tags = []string{"vip", "premium"}
func ParseMetadataValue(value string, target interface{}) error {
	// Try JSON unmarshal first (works for complex types)
	err := json.Unmarshal([]byte(value), target)
	if err == nil {
		return nil
	}

	// Handle simple types that might not unmarshal directly
	switch t := target.(type) {
	case *string:
		*t = value
		return nil
	case *int:
		val, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("cannot parse %q as int: %w", value, err)
		}
		*t = val
		return nil
	case *int64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %q as int64: %w", value, err)
		}
		*t = val
		return nil
	case *float64:
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %q as float64: %w", value, err)
		}
		*t = val
		return nil
	case *bool:
		val, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("cannot parse %q as bool: %w", value, err)
		}
		*t = val
		return nil
	}

	return fmt.Errorf("unsupported type or invalid JSON: %w", err)
}
