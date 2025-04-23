package json

import (
	"encoding/json"
	"os"
)

// WriteFile write data to JSON file
func WriteFile(filePath string, data any) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonBytes, 0664)
}

// WritePretty write pretty data to JSON file
func WritePretty(filePath string, data any) error {
	bs, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bs, 0664)
}

// ReadFile Read JSON file data
func ReadFile(filePath string, v any) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	return json.NewDecoder(file).Decode(v)
}
