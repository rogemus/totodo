package views

import (
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/ui"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type createProjectViewModel struct {
	focus        tui.Focus
	input        textinput.Model
	repo         repository.ProjectsRepository
	windowHeight int
	windowWidth  int
	isEditing    bool
	project      model.Project
}

func NewCreateProjectViewModel(repo repository.ProjectsRepository) createProjectViewModel {
	input := textinput.New()
	input.Placeholder = "Project name..."
	input.CharLimit = 125

	return createProjectViewModel{
		focus:     tui.CONFIRM_BTN,
		input:     input,
		repo:      repo,
		isEditing: false,
	}
}

func (m createProjectViewModel) Init() tea.Cmd { return nil }

func (m createProjectViewModel) View() string {
	cancelBtn := ui.CancelBtnStyle
	confirmBtn := ui.ConfirmBtnStyle
	confirmBtnText := "Create"
	titleText := "Create Project"

	if m.isEditing {
		confirmBtnText = "Update"
		titleText = "Update Project"
	}

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
				confirmBtn.Render(confirmBtnText),
			),
		),
	)

	dialogWrapper := lipgloss.JoinVertical(lipgloss.Top,
		ui.DialogTitleStyle.Render(titleText),
		ui.DialogBoxStyle.Render(dialogContent),
	)

	dialog := lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		dialogWrapper,
	)

	return dialog
}

func (m createProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := ui.WrapperStyle.GetFrameSize()
		m.windowWidth = msg.Width - h
		m.windowHeight = msg.Height - v

	case tui.ChangeViewMsg:
		m.project = model.Project{}
		m.isEditing = false
		m.focus = tui.NAME_INPUT
		m.input.SetValue("")
		m.input.Focus()

	case tui.ChangeViewWithProjectMsg:
		m.project = msg.Project
		m.isEditing = true
		m.focus = tui.NAME_INPUT
		m.input.SetValue(msg.Project.Name)
		m.input.Focus()

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

				if m.isEditing {
					m.project.Name = m.input.Value()
					m.repo.UpdateProject(m.project)
				} else {
					project := model.NewProject(m.input.Value())
					m.repo.CreateProject(project)
				}

				return m, tea.Batch(tui.NewChangeViewCmd(tui.PROJECTS_LIST_VIEW), tea.WindowSize())
			}

			if m.focus == tui.CANCEL_BTN {
				return m, tea.Batch(tui.NewChangeViewCmd(tui.PROJECTS_LIST_VIEW), tea.WindowSize())
			}

		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			return m, tea.Batch(tui.NewChangeViewCmd(tui.PROJECTS_LIST_VIEW), tea.WindowSize())
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
