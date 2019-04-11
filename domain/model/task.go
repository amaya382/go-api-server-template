package model

import "time"

type Task struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Title     string
	Contents  string
	Deadline  time.Time
	Completed bool

	TaskListID uint
}

type ArchivedTask struct {
	Task Task
}
