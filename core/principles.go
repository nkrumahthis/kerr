package core

import (
	"encoding/json"
	"os"
)

type Principle struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Score       int      `json:"score"`
}

func GetPrinciples() []Principle {
	data, err := os.ReadFile("principles.json")
	if err != nil {
		return nil
	}

	var principles []Principle
	err = json.Unmarshal(data, &principles)
	if err != nil {
		return nil
	}

	return principles
}

func SavePrinciples(principles []Principle) {

	data, _ := json.Marshal(principles)

	_ = os.WriteFile("principles.json", data, 0644)

}
