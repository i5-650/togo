package models

type Task struct {
	Title    string  `yaml:"title"`
	Deadline string  `yaml:"deadline"`
	Done     bool    `yaml:"done"`
	Subtasks []*Task `yaml:"subtasks,omitempty"`
}

type Todo struct {
	High   []*Task `yaml:"high"`
	Medium []*Task `yaml:"medium"`
	Low    []*Task `yaml:"low"`
}
