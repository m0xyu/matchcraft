package matcher

import (
	"encoding/json"
	"os"
)

type JSONEntity struct {
	ID         string             `json:"id"`
	Attributes map[string]float64 `json:"attributes"`
}

// これらのメソッドを実装することで Matchable インターフェースを満たします
func (j JSONEntity) GetID() string                     { return j.ID }
func (j JSONEntity) GetAttributes() map[string]float64 { return j.Attributes }

// LoadFromJSON は指定されたパスから JSON ファイルを読み込み、Matchable のスライスを返します
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
