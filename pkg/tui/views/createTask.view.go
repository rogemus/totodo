package views

import (
	"fmt"
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/ui"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type createTaskViewModel struct {
	focus        tui.Focus
	input        textinput.Model
	repo         repository.TasksRepository
	windowHeight int
	windowWidth  int
}

func NewCreateTaskViewModel(repo repository.TasksRepository) createTaskViewModel {
	input := textinput.New()
	input.Placeholder = "Task name..."
	input.CharLimit = 250

	return createTaskViewModel{
		focus: tui.CONFIRM_BTN,
		input: input,
		repo:  repo,
	}
}

func (m createTaskViewModel) Init() tea.Cmd { return nil }

func (m createTaskViewModel) View() string {
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
		// TODO: add focus styles
		m.input.View(),

		ui.DialogFooterStyle.Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				cancelBtn.Render("Cancel"),
				confirmBtn.Render("Create"),
			),
		),
	)

	dialogWrapper := lipgloss.JoinVertical(lipgloss.Top,
		ui.DialogTitleStyle.Render("Create Task"),
		ui.DialogBoxStyle.Render(dialogContent),
	)

	dialog := lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		dialogWrapper,
	)

	return dialog
}

func (m createTaskViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := ui.WrapperStyle.GetFrameSize()
		m.windowWidth = msg.Width - h
		m.windowHeight = msg.Height - v

	case tui.ChangeViewMsg:
		m.focus = tui.NAME_INPUT
		m.input.Focus()
		m.input.SetValue("")
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "tab":
			switch m.focus {
			case tui.NAME_INPUT:
				m.focus = tui.CONFIRM_BTN
			case tui.CONFIRM_BTN:
				m.focus = tui.CANCEL_BTN
			case tui.CANCEL_BTN:
				m.focus = tui.NAME_INPUT
			}

		case "enter":
			if m.focus == tui.CONFIRM_BTN {
				project := tui.State.SelectedProject
				task := model.NewTask(m.input.Value(), project.Id)

				_, err := m.repo.CreateTask(task)

				if err != nil {
					fmt.Printf("%v", err)
					return m, tea.Quit
				}

				return m, tea.Batch(tui.NewChangeViewCmd(tui.TASKS_LIST_VIEW), tea.WindowSize())
			}

			if m.focus == tui.CANCEL_BTN {
				return m, tea.Batch(tui.NewChangeViewCmd(tui.TASKS_LIST_VIEW), tea.WindowSize())
			}

		case "esc":
			return m, tea.Batch(tui.NewChangeViewCmd(tui.TASKS_LIST_VIEW), tea.WindowSize())
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
