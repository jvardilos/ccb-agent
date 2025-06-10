package fileio

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSON(data any, filePath string) error {
	// Marshal map to JSON
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Write to file
	err = os.WriteFile(filePath, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
