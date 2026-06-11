package main

import (
	"fmt"
	"github.com/KosyaGUT/taskmanager/internal/cli"
	"github.com/KosyaGUT/taskmanager/internal/task"
)

func main() {
	var tasks []task.Task
	var users []task.User
	var message string

	users = task.Hello(users)

	for flag := true; flag; {

		fmt.Println("Что вы выберите:\n" +
			"1. Создать задачу;\n" +
			"2. Удалить задачу;\n" +
			"3. Исправить задачу;\n" +
			"4. Показать список.\n" +
			"5. Профиль\n" +
			"6. Выход")
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
			task.Profile(users)
		case "6":
			flag = false
		}

	}

	fmt.Println("Удачи!")

}
