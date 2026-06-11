package main

import (
	"fmt"
)

func main() {
	var tasks []Task
	var message string

	for flag := true; flag; {

		fmt.Println("Что вы выберите:\n" +
			"1. Создать задачу;\n" +
			"2. Удалить задачу;\n" +
			"3. Исправить задачу;\n" +
			"4. Показать список.\n" +
			"5. Выход")
		message = acceptmessage()
		switch message {
		case "1":
			tasks = CreateTask(tasks)
		case "2":
			tasks, _ = DeleteTask(tasks)
		case "3":
			FixTask(tasks)
		case "4":
			AllTasks(tasks)
		case "5":
			flag = false
		}

	}

	fmt.Println("Удачи!")

}
