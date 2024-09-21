package enitities

import (
	"sync"
	"time"
)

var (
	instance *TaskManager
	once     sync.Once
)

type TaskManager struct {
	TaskInventory *TaskInventory
	UserInventory *UserInventory
	Reminders     []*Reminder
}

func GetInstance() *TaskManager {
	once.Do(func() {
		instance = &TaskManager{
			TaskInventory: NewTaskInventory(),
			UserInventory: NewUserInventory(),
		}
	})
	return instance
}

func (t *TaskManager) ChangeAssignee(taskId int, userId int) {
	oldUserId := t.TaskInventory.GetTaskById(taskId).AssignedTo
	t.TaskInventory.ChangeTaskAssignee(taskId, oldUserId, userId)
}

func (t *TaskManager) SetReminder(taskId int, time time.Time) {
	t.Reminders = append(t.Reminders, CreateReminder(taskId, time))
}
