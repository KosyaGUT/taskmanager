package main

import (
	"fmt"
	"github.com/KosyaGUT/taskmanager/internal/cli"
	"github.com/KosyaGUT/taskmanager/internal/task"
)

func main() {
	var tasks []task.Task
	var message string

	for flag := true; flag; {

		fmt.Println("Что вы выберите:\n" +
			"1. Создать задачу;\n" +
			"2. Удалить задачу;\n" +
			"3. Исправить задачу;\n" +
			"4. Показать список.\n" +
			"5. Выход")
		message = cli.Acceptmessage()
		switch message {
		case "1":
			tasks = task.CreateTask(tasks)
		case "2":
			tasks = task.DeleteTask(tasks)
		case "3":
			task.FixTask(tasks)
		case "4":
			task.AllTasks(tasks)
		case "5":
			flag = false
		}

	}

	fmt.Println("Удачи!")

}
