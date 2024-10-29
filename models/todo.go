package models

type Task struct {
	Title    string    `yaml:"title"`
	Deadline string    `yaml:"deadline"`
	Subtasks []Subtask `yaml:"subtasks,omitempty"`
}

type Subtask struct {
	Title    string `yaml:"title"`
	Deadline string `yaml:"deadline"`
}

type Todo struct {
	High   []Task `yaml:"high"`
	Medium []Task `yaml:"medium"`
	Low    []Task `yaml:"low"`
}
