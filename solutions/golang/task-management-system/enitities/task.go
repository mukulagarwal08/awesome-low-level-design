package enitities

import (
	"github.com/ashishps1/awesome-low-level-design/task-management-system/enums"
	"time"
)

type Task struct {
	id          int
	Title       string
	Description string
	Priority    enums.Priority
	Status      enums.Status
	CreatedAt   time.Time
	DueAt       time.Time
	AssignedTo  int
}

func NewTask(id int, title string, description string, priority enums.Priority, status enums.Status, dueTime time.Time, userId int) *Task {
	return &Task{
		id:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      status,
		CreatedAt:   time.Now(),
		DueAt:       dueTime,
		AssignedTo:  userId,
	}
}
func (t *Task) ChangeAssignedTo(userId int) {
	t.AssignedTo = userId
}

func (t *Task) GetAssignedTo() int {
	return t.AssignedTo
}
