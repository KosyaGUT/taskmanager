package task

import (
	"fmt"
	"github.com/KosyaGUT/taskmanager/internal/cli"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

func Registration(users []User) []User {
	fmt.Println("Добро пожаловать! Это меню регистрации. Напишите, пожалуйста, своё имя и фамилию")
	fistLastName := cli.Acceptmessage()
	fistlastnameSlice := strings.Split(fistLastName, " ")
	newUser := User{
		LastName:  fistlastnameSlice[0],
		FirstName: fistlastnameSlice[1],
	}
	fmt.Println("Отлично. Теперь напиши свой логин.")
	userLogin := cli.Acceptmessage()
	newUser.Login = userLogin
	fmt.Println("Отлично. Теперь напиши свой пароль.")
	password := cli.Acceptmessage()
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	hashedString := string(hashedBytes)
	newUser.Password = hashedString
	users = append(users, newUser)
	fmt.Printf("Вот ваш профиль\n"+
		"ID %v\n"+
		"Имя %v\n"+
		"Фамилия %v\n"+
		"Логин %v\n"+
		"Пароль %v\n", users[0].id, users[0].LastName, users[0].FirstName, users[0].Login, users[0].Password)
	fmt.Printf("Доброе пожаловать, %v %v\n", newUser.LastName, newUser.FirstName)
	return users
}

func Hello(users []User) []User {
	fmt.Println("Зравтвуйте! Вы уже зарегистрированы?\n" +
		"1. Да;\n" +
		"2. Нет.")
	message := cli.Acceptmessage()
	if message == "1" || message == "Да" {
		Authentification()
	} else {
		users = Registration(users)
	}
	return users
}

func Profile(users []User) {
	fmt.Printf("Вот ваш профиль\n"+
		"ID %v\n"+
		"Имя %v\n"+
		"Фамилия %v\n"+
		"Логин %v\n"+
		"Пароль %v\n", users[0].id, users[0].LastName, users[0].FirstName, users[0].Login, users[0].Password)
}

func Authentification() {
	return
}

func CreateTask(tasks []Task) []Task {
	for {
		newTask := Task{
			Title:       "",
			Description: "",
			Author:      "",
		}
		fmt.Println("Напишите заголовок задачи:")
		message := cli.Acceptmessage()
		if message == "стоп" {
			break
		}
		newTask.Title = message

		fmt.Println("Кого поставить исполнителем в задачу?")
		message = cli.Acceptmessage()
		newTask.Author = message

		fmt.Println("Если нужно, то можете добавить описание к задаче.")
		message = cli.Acceptmessage()
		newTask.Description = message

		tasks = append(
			tasks,
			newTask,
		)
		i := len(tasks) - 1
		fmt.Printf("ЗАДАЧА №%v БЫЛА ПРИНЯТА.\n"+
			"Заголовок: %v \n"+
			"Исполнитель: %v \n"+
			"Описание задачи: %v \n", i+1, tasks[i].Title, tasks[i].Author, tasks[i].Description)

		fmt.Println("Хотите ли еще задачу поставить?")

		message = cli.Acceptmessage()
		if strings.ToLower(message) == "да" || message == "" {
			continue
		} else {
			break
		}
	}
	return tasks
}

func DeleteTask(tasks []Task) []Task {
	fmt.Println("Какую задачу вы хотите удалить?")
	numTasks := ReadTasksNumber(tasks)

	fmt.Printf("Задача №%v %v будет удалена. Согласны? (да/нет)\n", numTasks, tasks[numTasks-1].Title)
	message := cli.Acceptmessage()
	if strings.ToLower(message) == "да" || message == "" {
		tasks = append(tasks[:numTasks-1], tasks[numTasks:]...)
	} else {
		fmt.Printf("Задача №%v не была удалена.", numTasks)
		return tasks
	}
	fmt.Printf("Задача №%v была удалена, нумерация обновлена.", numTasks)
	return tasks
}

func FixTask(tasks []Task) {
	fmt.Println("Какую задачу вы хотите исправить?")
	numTasks := ReadTasksNumber(tasks)

	fmt.Println("Что вы хотите исправить: \n" +
		"1. Заголовок \n" +
		"2. Описание \n" +
		"3. Исполнитель")
	message := cli.Acceptmessage()

	taskElement, _ := strconv.Atoi(message)
	if taskElement == 1 {
		fmt.Println("На что вы хотите поменять заголовок?")
		message = cli.Acceptmessage()
		tasks[numTasks-1].Title = message
	} else if taskElement == 2 {
		fmt.Println("На какие новое описание вы хотите поменять?")
		message = cli.Acceptmessage()
		tasks[numTasks-1].Description = message
	} else if taskElement == 3 {
		fmt.Println("На какого исполнителя хотите поменять?")
		message = cli.Acceptmessage()
		tasks[numTasks-1].Author = message
	}
	fmt.Printf("ИЗМЕНЕННАЯ ЗАДАЧА\n"+
		"Заголовок: %v \n"+
		"Исполнитель: %v \n"+
		"Описание задачи: %v \n", tasks[numTasks-1].Title, tasks[numTasks-1].Author, tasks[numTasks-1].Description)
}

func AllTasks(tasks []Task) {
	for j, task := range tasks {
		fmt.Printf("Задача №%v %v, Исполнитель: %v, Описание: %v\n", j+1, task.Title, task.Author, task.Description)
	}
}

func ReadTasksNumber(tasks []Task) int {
	for {
		message := cli.Acceptmessage()

		numTasks, err := strconv.Atoi(message)
		if err != nil {
			fmt.Println("Введите число")
			continue
		}

		if numTasks < 1 || numTasks > len(tasks) {
			fmt.Println("Такой задачи не существует!")
			continue
		}

		return numTasks
	}
}
