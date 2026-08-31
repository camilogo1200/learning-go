package task

import "time"

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
	Cancelled         = "Cancelled"
)

// transition table rules and lives with the type it governs
var allowedTransitions = map[Status][]Status{
	Todo:       {InProgress, Cancelled},
	InProgress: {Done, Cancelled, Todo},
	Done:       {},
	Cancelled:  {Todo},
}

type Priority int8

const (
	Urgent Priority = 1
	High            = 2
	Normal          = 3
	Low             = 4
)

type Task struct {
	Id          string
	ProjectId   string
	Title       string
	Description string
	Status      Status
	Priority    Priority
	AssigneeId  *string
	DueAt       *time.Time
	CreatedAt   time.time
	UpdatedAt   time.time
	Version     int
}
