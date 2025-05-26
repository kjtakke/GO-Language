package utils

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ToJSON converts any Go value (struct, map, slice) into a JSON-formatted byte slice.
func ToJSON(data interface{}) ([]byte, error) {
    return json.Marshal(data)
}

// ToJSONString converts any Go value into a formatted JSON string.
func ToJSONString(data interface{}) (string, error) {
    bytes, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// FromJSON parses a JSON byte slice into the provided destination (must be a pointer).
func FromJSON(jsonBytes []byte, target interface{}) error {
    if target == nil {
        return errors.New("target must be a non-nil pointer")
    }
    return json.Unmarshal(jsonBytes, target)
}

// FromJSONString parses a JSON string into the provided destination (must be a pointer).
func FromJSONString(jsonStr string, target interface{}) error {
    return FromJSON([]byte(jsonStr), target)
}

// FromJSONToMap decodes JSON bytes into a generic map[string]interface{}.
func FromJSONToMap(jsonBytes []byte) (map[string]interface{}, error) {
    var result map[string]interface{}
    err := json.Unmarshal(jsonBytes, &result)
    return result, err
}

// SafeGetString retrieves a string from a decoded JSON map.
func SafeGetString(data map[string]interface{}, key string) (string, bool) {
    val, ok := data[key]
    str, valid := val.(string)
    return str, ok && valid
}

// SafeGetNumber retrieves a float64 from a decoded JSON map.
func SafeGetNumber(data map[string]interface{}, key string) (float64, bool) {
    val, ok := data[key]
    num, valid := val.(float64)
    return num, ok && valid
}

// SafeGetBool retrieves a boolean from a decoded JSON map.
func SafeGetBool(data map[string]interface{}, key string) (bool, bool) {
    val, ok := data[key]
    b, valid := val.(bool)
    return b, ok && valid
}

// FlattenJSON extracts all keys and their values recursively from a nested map.
// Returns a slice of key-value pairs as [][2]string.
func FlattenJSON(data map[string]interface{}) [][2]string {
    var result [][2]string
    walkJSON(data, "", &result)
    return result
}

// walkJSON is a recursive helper for FlattenJSON.
func walkJSON(data map[string]interface{}, prefix string, acc *[][2]string) {
    for k, v := range data {
        fullKey := k
        if prefix != "" {
            fullKey = prefix + "." + k
        }

        switch val := v.(type) {
        case map[string]interface{}:
            walkJSON(val, fullKey, acc)
        default:
            *acc = append(*acc, [2]string{fullKey, fmt.Sprintf("%v", val)})
        }
    }
}
