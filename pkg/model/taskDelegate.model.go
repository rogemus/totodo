package model

import (
	"fmt"
	"io"
	"totodo/pkg/ui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type taskItemDelegate struct{}

func NewTaskItemDelegate() taskItemDelegate {
	return taskItemDelegate{}
}

func (d taskItemDelegate) Height() int { return 1 }

func (d taskItemDelegate) Spacing() int { return 0 }

func (d taskItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d taskItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	t, ok := listItem.(Task)

	if !ok {
		return
	}

	taskDone := t.Status == Status.DONE
	taskStatus := ui.MagentaTextStyle.Render("☐")
	taskName := ui.CyanTextStyle.Width(40).MaxWidth(80).Render(t.Name)
	taskIndex := ui.DimTextStyle.Render(fmt.Sprintf("%d.", index+1))

	if taskDone {
		taskStatus = ui.GreenTextStyle.Render("✓")
		taskName = ui.DimTextStyle.Width(40).MaxWidth(80).Strikethrough(true).Render(t.Name)
	}

	entry := ui.EntryStyle.Render(
		fmt.Sprintf("%s %s %s %s",
			taskIndex,
			taskStatus,
			taskName,
			ui.DimTextStyle.Render(fmt.Sprintf(" @%s ", t.GetTimeSinceCreation())),
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
