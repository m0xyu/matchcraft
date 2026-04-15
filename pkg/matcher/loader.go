package matcher

import (
	"encoding/json"
	"os"
)

// JSONEntity represents a data point loaded from a JSON file.
type JSONEntity struct {
	ID         string             `json:"id"`
	Attributes map[string]float64 `json:"attributes"`
}

// GetID returns the unique identifier of the JSON entity.
func (j JSONEntity) GetID() string { return j.ID }

// GetAttributes returns the attributes of the JSON entity as a map.
func (j JSONEntity) GetAttributes() map[string]float64 { return j.Attributes }

// LoadFromJSON loads a JSON file from the specified path and returns a slice of Matchable entities.
func LoadFromJSON(path string) ([]Matchable, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var entities []JSONEntity
	if err := json.Unmarshal(file, &entities); err != nil {
		return nil, err
	}

	results := make([]Matchable, len(entities))
	for i, v := range entities {
		results[i] = v
	}

	return results, nil
}
