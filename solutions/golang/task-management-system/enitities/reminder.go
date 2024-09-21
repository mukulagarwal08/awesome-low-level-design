package enitities

import "time"

type Reminder struct {
	TaskId int
	Time   time.Time
}

func CreateReminder(taskId int, t time.Time) *Reminder {
	return &Reminder{
		TaskId: taskId,
		Time:   t,
	}
}
