package views

import (
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type deleteTaskViewModel struct {
	focus        tui.Focus
	repo         repository.TasksRepository
	windowHeight int
	windowWidth  int
}

func NewDeleteTaskViewModel(repo repository.TasksRepository) deleteTaskViewModel {
	return deleteTaskViewModel{
		focus: tui.CONFIRM_BTN,
		repo:  repo,
	}
}

func (m deleteTaskViewModel) Init() tea.Cmd { return nil }

func (m deleteTaskViewModel) View() string {
	cancelBtn := ui.CancelBtnStyle
	confirmBtn := ui.ConfirmBtnStyle

	if m.focus == tui.CANCEL_BTN {
		cancelBtn = cancelBtn.Background(ui.BrightColors.Red).Bold(true)
	} else {
		cancelBtn = cancelBtn.Background(ui.NormalColors.Red).Bold(false)
	}

	if m.focus == tui.CONFIRM_BTN {
		confirmBtn = confirmBtn.Background(ui.BrightColors.Green).Bold(true)
	} else {
		confirmBtn = confirmBtn.Background(ui.NormalColors.Green).Bold(false)
	}

	dialogContent := lipgloss.JoinVertical(lipgloss.Top,
		ui.DialogTextStyle.Render("Are you sure you want to delete this task? This action cannot be undone."),

		ui.DialogFooterStyle.Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				cancelBtn.Render("Cancel"),
				confirmBtn.Render("Create"),
			),
		),
	)

	dialogWrapper := lipgloss.JoinVertical(lipgloss.Top,
		ui.DialogTitleStyle.Render("Delete Project?"),
		ui.DialogBoxStyle.Render(dialogContent),
	)

	dialog := lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		dialogWrapper,
	)

	return dialog

}

func (m deleteTaskViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := ui.WrapperStyle.GetFrameSize()
		m.windowWidth = msg.Width - h
		m.windowHeight = msg.Height - v

	case tui.ChangeViewMsg:
		m.focus = tui.CONFIRM_BTN
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "tab":
			switch m.focus {
			case tui.CANCEL_BTN:
				m.focus = tui.CONFIRM_BTN
			case tui.CONFIRM_BTN:
				m.focus = tui.CANCEL_BTN
			}

		case "enter":
			if m.focus == tui.CONFIRM_BTN {
				task := tui.State.SelectedTask
				m.repo.DeleteTask(task.Id)

				return m, tea.Batch(tui.NewChangeViewCmd(tui.TASKS_LIST_VIEW), tea.WindowSize())
			}

			return m, tea.Batch(tui.NewChangeViewCmd(tui.TASKS_LIST_VIEW), tea.WindowSize())

		case "esc":
			return m, tea.Batch(tui.NewChangeViewCmd(tui.TASKS_LIST_VIEW), tea.WindowSize())
		}
	}

	return m, cmd
}
