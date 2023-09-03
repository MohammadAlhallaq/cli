package todo

import (
	"encoding/json"
	"os"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = os.WriteFile(filename, b, 0644)

	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	d, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	err = json.Unmarshal(d, &items)

	if err != nil {
		return []Item{}, err
	}
	return items, nil
}
