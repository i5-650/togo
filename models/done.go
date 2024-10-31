package models

import "time"

type DoneTask struct {
	Title      string    `yaml:"title"`
	DeletedAt  time.Time `yaml:"deleted_at"`
	Importance string    `yaml:"importance"`
	Subtasks   []*Task   `yaml:"subtasks,omitempty"`
}

type Done struct {
	Tasks []DoneTask `yaml:"tasks"`
}
