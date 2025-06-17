package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveToJSON a given data structure object to file at location. Validates filepath given
func SaveToJSON[T any](data []T, filepath string) error {

	if !ValidateFilepath(filepath) {
		return &StringError{filepath, FilepathError}
	}

	m, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("Error marshaling data: %s", err)
	}

	err = os.WriteFile(filepath, m, 0644)
	if err != nil {
		return fmt.Errorf("Error writing data to fole: %s", err)
	}

	return nil
}

// LoadFromJSON loads data from a JSON f0]\ile into a struct
func LoadFromJSON[T any](filepath string) ([]T, error) {
	if !ValidateFilepath(filepath) {
		return nil, &StringError{filepath, FilepathError}
	}
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %s", err)
	}

	var data []T
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal data: %s", err)
	}

	return data, nil
}
