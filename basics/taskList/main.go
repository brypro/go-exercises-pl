package main

import (
	"fmt"
)

type task struct {
	name     string
	desc     string
	complete bool
}

func (t *task) MarkAsComplete() {
	t.complete = true
}
func (t *task) UpdateDescription(newDesc string) {
	t.desc = newDesc
}
func (t *task) UpdateName(newName string) {
	t.name = newName
}
func NewTask(name, desc string) *task {
	t := task{name, desc, false}
	return &t
}
func main() {
	var list taskList
	list.AddTask("study", "study all my courses")
	list.AddTask("read", "task 2")
	list.AddTask("play games", "task 3")
	list.AddTask("play 2", "task 4")
	list.AddTask("play 3", "task 5")
	list.String()
	list.DeleteTask(1) //segundo
	list.String()
	list.CompleteTask(0)
	list.CompleteTask(3)
	list.StringCompletes()
}

type taskList struct {
	tasks []*task
}

func (tl *taskList) CompleteTask(index int) {
	tl.tasks[index].MarkAsComplete()
}
func (tl *taskList) DeleteTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}
func (tl *taskList) AddTask(name, desc string) {
	tl.tasks = append(tl.tasks, NewTask(name, desc))
}
func (tl *taskList) String() {
	fmt.Printf("********** TASK LIST  *************\n")
	for i, task := range tl.tasks {
		fmt.Printf("[%d] Name: %s, Desc: %s completed: %v\n", i, task.name, task.desc, task.complete)
	}
	fmt.Printf("***********************************\n\n\n")
}

func (tl *taskList) StringCompletes() {
	comp := fmt.Sprintf("********** COMPLETE TASK *************\n")
	incom := fmt.Sprintf("********* INCOMPLETE TASK ************\n")
	for i, task := range tl.tasks {
		if task.complete {
			comp += fmt.Sprintf("[%d] Name: %s, Desc: %s completed: %v\n", i, task.name, task.desc, task.complete)
		} else {

			incom += fmt.Sprintf("[%d] Name: %s, Desc: %s completed: %v\n", i, task.name, task.desc, task.complete)
		}
	}
	comp += fmt.Sprintf("***********************************\n\n")
	incom += fmt.Sprintf("***********************************\n\n")
	fmt.Println(comp, incom)
}
