package views

import (
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type tasksListViewModel struct {
	list list.Model
	done bool
	repo repository.TasksRepository
}

func NewTasksListViewModel(repo repository.TasksRepository) tasksListViewModel {
	m := tasksListViewModel{
		repo: repo,
		list: list.New([]list.Item{
			model.Task{Name: "Task 1"},
		}, list.NewDefaultDelegate(), 10, 10),
	}

	return m
}

func (m tasksListViewModel) Init() tea.Cmd {
	return nil
}

func (m tasksListViewModel) View() string {
	return m.list.View()
}

func (m tasksListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tui.ChangeToTasksListViewMsg:
		tasks, _ := m.repo.GetTasks(tui.State.SelectedTask.Id)
		m.list.SetItems(convertToListitem(tasks))

		return m, nil
	}

	return m, cmd
}

func convertToListitem(tasks []model.Task) []list.Item {
	var items []list.Item

	for _, t := range tasks {
		items = append(items, list.Item(t))
	}

	return items
}
