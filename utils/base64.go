package utils

import (
    "encoding/base64"
    "encoding/json"
    "errors"
)

// EncodeStringToBase64 returns the Base64-encoded representation of a string.
func EncodeStringToBase64(input string) string {
    return base64.StdEncoding.EncodeToString([]byte(input))
}

// DecodeBase64ToString decodes a Base64-encoded string into a plain string.
func DecodeBase64ToString(encoded string) (string, error) {
    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return "", err
    }
    return string(decodedBytes), nil
}

// EncodeJSONToBase64 takes a Go map or struct, marshals it to JSON, and encodes it in Base64.
func EncodeJSONToBase64(data interface{}) (string, error) {
    jsonBytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

// DecodeBase64JSONToMap decodes a Base64 string assumed to contain JSON and parses it into a map.
func DecodeBase64JSONToMap(encoded string) (map[string]interface{}, error) {
    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return nil, err
    }

    var result map[string]interface{}
    if err := json.Unmarshal(decodedBytes, &result); err != nil {
        return nil, err
    }

    return result, nil
}

// DecodeBase64JSONToStruct decodes a Base64 JSON string into a given struct pointer.
func DecodeBase64JSONToStruct(encoded string, target interface{}) error {
    if target == nil {
        return errors.New("target must be a non-nil pointer to a struct")
    }

    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return err
    }

    return json.Unmarshal(decodedBytes, target)
}
