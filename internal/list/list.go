package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-cli/internal/status"
)

// ListTask prints tasks from task.json filtered by status option.
//
// option values:
//   - 1: todo
//   - 2: in-progress
//   - 3: done
//   - 4: all
//
// For each matched task, it prints description + status, then ID, UPDATEDAT, and CREATEDAT.
func ListTask(option int)error{
	_,err:=os.Stat("task.json")
	if os.IsNotExist(err){
		return errors.New("task.json is not exist")
	}else if err!=nil{
		return err
	}
	var tasks []status.Task
	taskText,err:=os.ReadFile("task.json")
	if err!=nil{
		return err
	}
	if len(taskText)==0{
		return nil 
	}
	err=json.Unmarshal(taskText,&tasks)
	if err != nil{
		return err
	}
	mapTask:=make(map[string]int)
	mapTask["todo"]=1
	mapTask["in-progress"]=2
	mapTask["done"]=3
	for _,task:=range tasks{
		if option==4||option==mapTask[task.STATUS]{
			fmt.Println(task.DESCRIPTION,task.STATUS)
			fmt.Printf("ID: %v\n",task.ID)
			fmt.Printf("Updated at:%v\n",task.UPDATEDAT)
			fmt.Printf("Created at:%v\n",task.CREATEDAT)
			fmt.Printf("\n")
		}
	}
	return nil
}