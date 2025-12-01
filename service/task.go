package service

import (
	"errors"
	"mini-project-2/model"
	"mini-project-2/utils"
	"strings"
)

// function add task
func AddTask(task, status, priority string) error {
	tasks, _ := utils.LoadTasks()

	for _, t := range tasks {
		if t.Task == task {
			return errors.New("task duplikat")
		}
	}

	taskp := model.Task{
	ID:          len(tasks) + 1,
	Task:       task,
	Status:      status,
	Priority:    priority,
}
	tasks = append(tasks, taskp)

	return utils.SaveTasks(tasks)
}

// function list task
func ListTasks() ([]model.Task, error) {
	return utils.LoadTasks()
}

// function mark task
func MarkDone(id int) error {
	tasks, err := utils.LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = "completed" 
			found = true
			break
		}
	}

	if !found {
		return errors.New("task tidak ditemukan")
	}

	return utils.SaveTasks(tasks) 
}

// function delete task
func DeleteTask(id int) error {
	tasks, _ := utils.LoadTasks()
	newTasks := []model.Task{}

	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		}
	}
	return utils.SaveTasks(newTasks)
}

// function search task
func SearchTask(keyword string) ([]model.Task, error) {
	tasks, _ := utils.LoadTasks()
	var result []model.Task
	keyword = strings.ToLower(keyword)

	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Task), keyword) {
			result = append(result, t)
		}
	}
	return result, nil
}
