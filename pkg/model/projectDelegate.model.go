package model

import (
	"fmt"
	"io"
	"totodo/pkg/ui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type projectItemDelegate struct{}

func NewProjectItemDelegate() projectItemDelegate {
	return projectItemDelegate{}
}

func (d projectItemDelegate) Height() int { return 1 }

func (d projectItemDelegate) Spacing() int { return 0 }

func (d projectItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d projectItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	p, ok := listItem.(Project)

	if !ok {
		return
	}

	entry := ui.EntryStyle.Render(
		fmt.Sprintf("%s %s %s",
			ui.CyanTextStyle.Width(25).MaxWidth(50).Render(p.Name),
			ui.DimTextStyle.Render(fmt.Sprintf(" [%d/%d] ", p.TasksDoneCount, p.TasksCount)),
			ui.DimTextStyle.Render(fmt.Sprintf(" @%s ", p.GetTimeSinceCreation())),
		),
	)

	indicator := ui.GreenTextStyle.Render("|")

	fn := func() string {
		if index == m.Index() {
			return fmt.Sprintf("%s %s", indicator, entry)
		}

		return fmt.Sprintf("  %s", entry)
	}

	fmt.Fprint(w, fn())
}
