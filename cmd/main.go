package main

import (
	"errors"
	"os"
	"strconv"
	"task-cli/internal/errorPrint"
	"task-cli/internal/help"
	"task-cli/internal/list"
	"task-cli/internal/status"
)

func main() {
	if len(os.Args)<2{
		help.Help()
	}
	switch os.Args[1]{
	case "help":
		help.Help()
	case "add":
		if len(os.Args)>4 {
			errorPrint.ErrPrint(errors.New("Too many arguments"))
		}else if len(os.Args)<=2{
			errorPrint.ErrPrint(errors.New("Missing arguments"))
		}else if len(os.Args)==3{
			var task status.Task
			task.DESCRIPTION=os.Args[2]
			task.STATUS="todo"
			err:=status.AddTask(task)
			if err != nil{
				errorPrint.ErrPrint(err)
			}
		}else{
			var task status.Task
			task.DESCRIPTION=os.Args[2]
			task.STATUS=os.Args[3]
			err:=status.AddTask(task)
			if err != nil{
				errorPrint.ErrPrint(err)
			}
		}
	case "update":
		if len(os.Args)>4 {
			errorPrint.ErrPrint(errors.New("Too many arguments"))
		}else if len(os.Args)<4{
			errorPrint.ErrPrint(errors.New("Missing arguments"))
		}else{
			id,err:=strconv.Atoi(os.Args[2])
			if err != nil{
				errorPrint.ErrPrint(err)
			}
			err=status.UpdateTask(id,"description",os.Args[3])
			if err != nil{
				errorPrint.ErrPrint(err)
			}
		}
	case "mark":
		if len(os.Args)>4 {
			errorPrint.ErrPrint(errors.New("Too many arguments"))
		}else if len(os.Args)<4{
			errorPrint.ErrPrint(errors.New("Missing arguments"))
		}else{
			id,err:=strconv.Atoi(os.Args[2])
			if err != nil{
				errorPrint.ErrPrint(err)
			}
			err=status.UpdateTask(id,"status",os.Args[3])
			if err != nil{
				errorPrint.ErrPrint(err)
			}
		}
	case "delete":
		if len(os.Args)>3 {
			errorPrint.ErrPrint(errors.New("Too many arguments"))
		}else if len(os.Args)<3{
			errorPrint.ErrPrint(errors.New("Missing arguments"))
		}else{
			id,err:=strconv.Atoi(os.Args[2])
			if err != nil{
				errorPrint.ErrPrint(err)
			}
			err=status.DeleteTask(id)
			if err != nil{
				errorPrint.ErrPrint(err)
			}
		}
	case "list":
		if len(os.Args)>4 {
			errorPrint.ErrPrint(errors.New("Too many arguments"))
		}else if len(os.Args)==2{
			err:=list.ListTask(4)
			if err!=nil{
				errorPrint.ErrPrint(err)
			}
		}else{
			switch os.Args[2]{
			case "todo":
				err:=list.ListTask(1)
				if err!=nil{
					errorPrint.ErrPrint(err)
				}
			case "in-progress":
				err:=list.ListTask(2)
				if err!=nil{
					errorPrint.ErrPrint(err)
				}
			case "done":
				err:=list.ListTask(3)
				if err!=nil{
					errorPrint.ErrPrint(err)
				}
			case "all":
				err:=list.ListTask(4)
				if err!=nil{
					errorPrint.ErrPrint(err)
				}
			default:
				errorPrint.ErrPrint(errors.New("Unknown argument"))
			}
		}
	default:
		errorPrint.ErrPrint(errors.New("Unknown command"))
	}
}
