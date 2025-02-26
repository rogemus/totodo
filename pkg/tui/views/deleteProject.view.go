package views

import (
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type deleteProjectViewModel struct {
	focus        tui.Focus
	project      model.Project
	repo         repository.ProjectsRepository
	windowHeight int
	windowWidth  int
}

func NewDeleteProjectViewModel(repo repository.ProjectsRepository) deleteProjectViewModel {
	return deleteProjectViewModel{
		focus: tui.CONFIRM_BTN,
		repo:  repo,
	}
}

func (m deleteProjectViewModel) Init() tea.Cmd { return nil }

func (m deleteProjectViewModel) View() string {
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
		ui.DialogTextStyle.Render("Are you sure you want to delete this project? This action cannot be undone."),

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

func (m deleteProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := ui.WrapperStyle.GetFrameSize()
		m.windowWidth = msg.Width - h
		m.windowHeight = msg.Height - v

	case tui.ChangeViewWithProjectMsg:
		m.focus = tui.CONFIRM_BTN
		m.project = msg.Project

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "tab":
			switch m.focus {
			case tui.CONFIRM_BTN:
				m.focus = tui.CANCEL_BTN
			case tui.CANCEL_BTN:
				m.focus = tui.CONFIRM_BTN
			}

		case "enter":
			if m.focus == tui.CONFIRM_BTN {
				m.repo.DeleteProject(m.project.Id)
				return m, tea.Batch(tui.NewChangeViewCmd(tui.PROJECTS_LIST_VIEW), tea.WindowSize())
			}

			return m, tea.Batch(tui.NewChangeViewCmd(tui.PROJECTS_LIST_VIEW), tea.WindowSize())

		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			return m, tea.Batch(tui.NewChangeViewCmd(tui.PROJECTS_LIST_VIEW), tea.WindowSize())
		}
	}

	return m, cmd
}
