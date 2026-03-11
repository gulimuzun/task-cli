package status

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Task represents a single task record.
// It is serialized/deserialized to/from task.json: ID is the unique identifier,
// DESCRIPTION is the task text, STATUS is the task state, and
// CREATEDAT/UPDATEDAT store the created/updated timestamps (JSON: createdAt/updatedAt).
type Task struct {
	ID          int    `json:"id"`
	DESCRIPTION string `json:"description"`
	STATUS      string `json:"status"`
	CREATEDAT   string `json:"createdAt"`
	UPDATEDAT   string `json:"updatedAt"`
}

// AddTask appends a new task to task.json.
// If task.json does not exist, it creates the file first. It assigns an auto-incremented ID,
// sets CREATEDAT and UPDATEDAT to the current time, then writes the updated task list back.
func AddTask(newTask Task) error {
	_, err := os.Stat("task.json")
	if os.IsNotExist(err) {
		f, err := os.Create("task.json")
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Println("Successfully create task.json")
	} else if err != nil {
		return err
	}
	var tasks []Task
	taskText, err := os.ReadFile("task.json")
	if err !=nil{
		return err
	}
	if len(taskText) == 0 {
		tasks = []Task{}
	}else {
		err = json.Unmarshal(taskText, &tasks)
	}
	if err != nil {
		return err
	}
	newID := 1
	if len(tasks) != 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}
	newTask.ID = newID
	newTask.CREATEDAT = time.Now().Format("2006-01-02 15:04:05")
	newTask.UPDATEDAT = newTask.CREATEDAT
	tasks = append(tasks, newTask)
	newTasksText, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("task.json", newTasksText, 0644)
	fmt.Println("Successfully add new task,ID:",newID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTask updates a task in task.json by its ID.
// Supported options are "status" and "description". When updating status, the allowed values are
// "todo", "in-progress", and "done". On success it refreshes UPDATEDAT and persists the changes
// back to task.json.
func UpdateTask(taskID int, option string, newContent string) error {
	_, err := os.Stat("task.json")
	if os.IsNotExist(err) {
		f, err := os.Create("task.json")
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Println("Successfully create task.json")
	} else if err != nil {
		return err
	}
	var tasks []Task
	taskText, err := os.ReadFile("task.json")
	if err != nil {
		return err
	}
	if len(taskText) == 0 {
		return errors.New("Can`t find the task for the id ")
	} else {
		err = json.Unmarshal(taskText, &tasks)
	}
	if err != nil {
		return err
	}
	isTaskExist:=false
	for index, task := range tasks {
		if task.ID == taskID {
			isTaskExist=true
			if strings.EqualFold("status", option) {
				if strings.EqualFold("todo", newContent) || strings.EqualFold("in-progress", newContent) || strings.EqualFold("done", newContent) {
					tasks[index].STATUS = newContent
					tasks[index].UPDATEDAT=time.Now().Format("2006-01-02 15:04:05")
				} else {
					return errors.New("invalid status")
				}
			} else if strings.EqualFold("description", option) {
				tasks[index].DESCRIPTION = newContent
				tasks[index].UPDATEDAT=time.Now().Format("2006-01-02 15:04:05")
			} else {
				return errors.New("Option not exist")
			}
			break
		}
	}
	if !isTaskExist{
		return errors.New("Can`t find the task for the id")	
	}
	newTaskText,err:=json.MarshalIndent(tasks,""," ")
	if err != nil{
		return err
	}
	err=os.WriteFile("task.json",newTaskText,0644)
	if err != nil{
		return err
	}
	fmt.Println("Successfully update the task")
	return nil
}
// DeleteTask removes a task from task.json by its ID.
// It rewrites task.json without the matching task and returns an error if the ID does not exist.
func DeleteTask(taskID int) error{
	_, err := os.Stat("task.json")
	if os.IsNotExist(err) {
		f, err := os.Create("task.json")
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Println("Successfully create task.json")
	} else if err != nil {
		return err
	}
	var tasks []Task
	taskText, err := os.ReadFile("task.json")
	if err != nil {
		return err
	}
	if len(taskText) == 0 {
		return errors.New("Can`t find the task for the id ")
	} else {
		err = json.Unmarshal(taskText, &tasks)
	}
	if err != nil {
		return err
	}
	var newTasks []Task
	newTasks=[]Task{}
	isDelete:=false
  	for _,task := range tasks{
		if task.ID==taskID{
			isDelete=true
			continue
		}
		newTasks=append(newTasks, task)
	}
	if !isDelete{
		return errors.New("Can`t find the task for the ID")
	}
	newTaskText,err:=json.MarshalIndent(newTasks,""," ")
	if err != nil{
		return err
	}
	err=os.WriteFile("task.json",newTaskText,0644)
	if err != nil{
		return err
	}
	fmt.Println("Successfully delete the task")
	return nil
}