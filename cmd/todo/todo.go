package todo

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	Position int
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

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}

func (i *Item) PrettyPri() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 2 {
		return "(2)"
	}
	if i.Priority == 3 {
		return "(3)"
	}

	return " "
}

func (i *Item) SetLabel(label int) {
	i.Position = label
}
