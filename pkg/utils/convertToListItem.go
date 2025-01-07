package utils

import (
	"totodo/pkg/model"

	"github.com/charmbracelet/bubbles/list"
)

type listItem interface {
	model.Project | model.Task
}

func ConvertToListitem[T listItem](tasks []T) []list.Item {
	var items []list.Item

	for _, t := range tasks {
		items = append(items, list.Item(t))
	}

	return items
}
