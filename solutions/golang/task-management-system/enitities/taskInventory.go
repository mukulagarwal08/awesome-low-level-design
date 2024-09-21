package enitities

import (
	"github.com/ashishps1/awesome-low-level-design/task-management-system/enums"
	"time"
)

type TaskInventory struct {
	UserTasks map[int]map[enums.Priority]map[enums.Status]map[int]struct{}
	TasksData map[int]*Task
}

func NewTaskInventory() *TaskInventory {
	return &TaskInventory{
		UserTasks: make(map[int]map[enums.Priority]map[enums.Status]map[int]struct{}),
		TasksData: make(map[int]*Task),
	}
}

func (t *TaskInventory) AddTask(id int, title, description string, priority enums.Priority, status enums.Status, dueTime time.Time, userIDAssigned int) {
	task := NewTask(id, title, description, priority, status, dueTime, userIDAssigned)
	t.TasksData[id] = task

}
func (t *TaskInventory) ChangeTaskAssignee(taskId int, oldUserId int, newUserId int) {
	task := t.TasksData[taskId]
	oldUserTaskData := t.UserTasks[oldUserId][task.Priority][task.Status]
	//remove task from old user
	delete(oldUserTaskData, taskId)

	//add task to new user
	t.UserTasks[newUserId][task.Priority][task.Status][taskId] = struct{}{}
}

func (t *TaskInventory) GetTaskById(taskId int) *Task {
	return t.TasksData[taskId]
}

func (t *TaskInventory) GetTaskByAssignedUserId(userId int) []*Task {
	var result []*Task

	for _, priorityMapItem := range t.UserTasks[userId] {
		for _, statusMapItem := range priorityMapItem {
			for taskId, _ := range statusMapItem {
				result = append(result, t.GetTaskById(taskId))
			}
		}
	}
	return result
}

func (t *TaskInventory) GetTaskByPriority(priority enums.Priority) []*Task {
	var result []*Task

	for _, userMapItem := range t.UserTasks {
		for _, statusMapItem := range userMapItem[priority] {
			for taskId, _ := range statusMapItem {
				result = append(result, t.GetTaskById(taskId))
			}
		}
	}
	return result
}
