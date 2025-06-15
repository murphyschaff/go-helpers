package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveToJSON a given data structure object to file at location. Validates filepath given
func SaveToJSON(filepath string, data *any) error {

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

// LoadFromJSON loads data from a JSON file into a struct
func LoadFromJSON(filepath string, data ...*any) error {
	if !ValidateFilepath(filepath) {
		return &StringError{filepath, FilepathError}
	}
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Unable to read file: %s", err)
	}

	err = json.Unmarshal(fileData, data)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal data: %s", err)
	}

	return nil
}
