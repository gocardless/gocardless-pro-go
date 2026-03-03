package gocardless

import (
	"testing"
)

func TestToMetadataValue(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		// Strings
		{"string unchanged", "hello", "hello"},
		{"empty string", "", ""},

		// Integers
		{"int", 12345, "12345"},
		{"int zero", 0, "0"},
		{"int negative", -42, "-42"},
		{"int8", int8(127), "127"},
		{"int16", int16(32767), "32767"},
		{"int32", int32(2147483647), "2147483647"},
		{"int64", int64(9223372036854775807), "9223372036854775807"},

		// Unsigned integers
		{"uint", uint(12345), "12345"},
		{"uint8", uint8(255), "255"},
		{"uint16", uint16(65535), "65535"},
		{"uint32", uint32(4294967295), "4294967295"},
		{"uint64", uint64(18446744073709551615), "18446744073709551615"},

		// Floats
		{"float32", float32(3.14), "3.14"},
		{"float64", float64(3.14159), "3.14159"},
		{"float zero", float64(0), "0"},
		{"float negative", float64(-2.5), "-2.5"},

		// Booleans
		{"bool true", true, "true"},
		{"bool false", false, "false"},

		// Nil
		{"nil", nil, "null"},

		// Slices (JSON serialization)
		{"string slice", []string{"vip", "premium"}, `["vip","premium"]`},
		{"int slice", []int{1, 2, 3}, `[1,2,3]`},
		{"empty slice", []string{}, `[]`},

		// Maps (JSON serialization)
		{"string map", map[string]string{"theme": "dark"}, `{"theme":"dark"}`},
		{"mixed map", map[string]interface{}{"count": 5, "enabled": true}, `{"count":5,"enabled":true}`},
		{"empty map", map[string]string{}, `{}`},

		// Structs (JSON serialization)
		{"struct", struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{"John", 30}, `{"name":"John","age":30}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMetadataValue(tt.input)
			if result != tt.expected {
				t.Errorf("ToMetadataValue(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]string
	}{
		{
			name: "mixed types",
			input: map[string]interface{}{
				"user_id":   12345,
				"is_active": true,
				"name":      "John",
			},
			expected: map[string]string{
				"user_id":   "12345",
				"is_active": "true",
				"name":      "John",
			},
		},
		{
			name: "with arrays and objects",
			input: map[string]interface{}{
				"tags":  []string{"vip", "premium"},
				"prefs": map[string]string{"theme": "dark"},
			},
			expected: map[string]string{
				"tags":  `["vip","premium"]`,
				"prefs": `{"theme":"dark"}`,
			},
		},
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: map[string]string{},
		},
		{
			name: "already strings",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "nil values",
			input: map[string]interface{}{
				"nullable": nil,
			},
			expected: map[string]string{
				"nullable": "null",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMetadata(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("ToMetadata() length = %d, want %d", len(result), len(tt.expected))
			}

			for key, expectedValue := range tt.expected {
				if actualValue, ok := result[key]; !ok {
					t.Errorf("ToMetadata() missing key %q", key)
				} else if actualValue != expectedValue {
					t.Errorf("ToMetadata()[%q] = %q, want %q", key, actualValue, expectedValue)
				}
			}
		})
	}
}

func TestIsValidMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected bool
	}{
		{
			name: "valid - all strings",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			expected: true,
		},
		{
			name: "invalid - contains int",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": 12345,
			},
			expected: false,
		},
		{
			name: "invalid - contains bool",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": true,
			},
			expected: false,
		},
		{
			name: "invalid - contains map",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": map[string]string{"nested": "value"},
			},
			expected: false,
		},
		{
			name:     "valid - empty map",
			input:    map[string]interface{}{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidMetadata(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidMetadata() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseMetadataValue(t *testing.T) {
	t.Run("parse to string", func(t *testing.T) {
		var result string
		err := ParseMetadataValue("hello", &result)
		if err != nil {
			t.Errorf("ParseMetadataValue() error = %v", err)
		}
		if result != "hello" {
			t.Errorf("ParseMetadataValue() = %q, want %q", result, "hello")
		}
	})

	t.Run("parse to int", func(t *testing.T) {
		var result int
		err := ParseMetadataValue("12345", &result)
		if err != nil {
			t.Errorf("ParseMetadataValue() error = %v", err)
		}
		if result != 12345 {
			t.Errorf("ParseMetadataValue() = %d, want %d", result, 12345)
		}
	})

	t.Run("parse to int64", func(t *testing.T) {
		var result int64
		err := ParseMetadataValue("9223372036854775807", &result)
		if err != nil {
			t.Errorf("ParseMetadataValue() error = %v", err)
		}
		if result != 9223372036854775807 {
			t.Errorf("ParseMetadataValue() = %d, want %d", result, int64(9223372036854775807))
		}
	})

	t.Run("parse to float64", func(t *testing.T) {
		var result float64
		err := ParseMetadataValue("3.14", &result)
		if err != nil {
			t.Errorf("ParseMetadataValue() error = %v", err)
		}
		if result != 3.14 {
			t.Errorf("ParseMetadataValue() = %f, want %f", result, 3.14)
		}
	})

	t.Run("parse to bool", func(t *testing.T) {
		var result bool
		err := ParseMetadataValue("true", &result)
		if err != nil {
			t.Errorf("ParseMetadataValue() error = %v", err)
		}
		if result != true {
			t.Errorf("ParseMetadataValue() = %v, want %v", result, true)
		}
	})

	t.Run("parse to slice", func(t *testing.T) {
		var result []string
		err := ParseMetadataValue(`["vip","premium"]`, &result)
		if err != nil {
			t.Errorf("ParseMetadataValue() error = %v", err)
		}
		if len(result) != 2 || result[0] != "vip" || result[1] != "premium" {
			t.Errorf("ParseMetadataValue() = %v, want [vip premium]", result)
		}
	})

	t.Run("parse to map", func(t *testing.T) {
		var result map[string]string
		err := ParseMetadataValue(`{"theme":"dark"}`, &result)
		if err != nil {
			t.Errorf("ParseMetadataValue() error = %v", err)
		}
		if result["theme"] != "dark" {
			t.Errorf("ParseMetadataValue() = %v, want map[theme:dark]", result)
		}
	})

	t.Run("parse invalid int", func(t *testing.T) {
		var result int
		err := ParseMetadataValue("not-a-number", &result)
		if err == nil {
			t.Error("ParseMetadataValue() expected error for invalid int")
		}
	})

	t.Run("parse invalid bool", func(t *testing.T) {
		var result bool
		err := ParseMetadataValue("not-a-bool", &result)
		if err == nil {
			t.Error("ParseMetadataValue() expected error for invalid bool")
		}
	})

	t.Run("parse invalid JSON", func(t *testing.T) {
		var result []string
		err := ParseMetadataValue("not valid json", &result)
		if err == nil {
			t.Error("ParseMetadataValue() expected error for invalid JSON")
		}
	})
}
