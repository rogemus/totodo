package pkg

type Task struct {
	Description string
	Tag         string
	Project     string
}

func NewTask(desc, tag, proj string) Task {
	return Task{
		Description: desc,
		Tag:         tag,
		Project:     proj,
	}
}
