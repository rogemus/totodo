package views

import (
	"fmt"
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
)

type tasksListViewModel struct {
	project *model.Project
}

func NewTasksListViewModel(repo repository.TasksRepository) tasksListViewModel {
	return tasksListViewModel{}
}

func (m tasksListViewModel) Init() tea.Cmd {
	return nil
}

func (m tasksListViewModel) View() string {
	return fmt.Sprintf("taskListViewModel, %s", tui.State.SelectedProject.Name)
}

func (m tasksListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}
