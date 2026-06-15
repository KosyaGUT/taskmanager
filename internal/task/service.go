package task

import (
	"fmt"
	"github.com/KosyaGUT/taskmanager/internal/cli"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

func Registration(users []User) []User {
	fmt.Println("Добро пожаловать! Это меню регистрации. Напишите, пожалуйста, своё имя и фамилию")
	fistLastName := cli.Acceptmessage()
	fistlastnameSlice := strings.Split(fistLastName, " ")
	newUser := User{
		Id:        uuid.New(),
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
	err := SaveUsers(users)
	if err != nil {
		return nil
	}
	fmt.Printf("Доброе пожаловать, %v %v\n", newUser.LastName, newUser.FirstName)
	return users
}

func Hello(users []User) ([]User, *User) {

	fmt.Println("Здравствуйте! Вы уже зарегистрированы?")
	fmt.Println("1. Да")
	fmt.Println("2. Нет")

	message := cli.Acceptmessage()

	if message == "1" {

		user, err := Authentication(users)

		if err != nil {
			fmt.Println(err)
			return users, nil
		}

		return users, user
	}

	users = Registration(users)

	return users, &users[len(users)-1]
}

func Profile(user *User) {
	fmt.Printf(
		"ID: %v\n"+
			"Имя: %v\n"+
			"Фамилия: %v\n"+
			"Логин: %v\n",
		user.Id,
		user.FirstName,
		user.LastName,
		user.Login,
	)
}

func Authentication(users []User) (*User, error) {
	fmt.Print("Логин: ")
	login := cli.Acceptmessage()

	fmt.Print("Пароль: ")
	password := cli.Acceptmessage()

	for i := range users {

		if users[i].Login != login {
			continue
		}

		err := bcrypt.CompareHashAndPassword(
			[]byte(users[i].Password),
			[]byte(password),
		)

		if err == nil {
			return &users[i], nil
		}
	}

	return nil, fmt.Errorf("неверный логин или пароль")
}

func CreateTask(tasks []Task, currentUser *User) []Task {
	for {
		newTask := Task{
			Id:           uuid.New(),
			CreatorLogin: currentUser.Login,
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
	err := SaveTasks(tasks)
	if err != nil {
		return nil
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
		fmt.Printf(
			"Задача №%v %v, Создал: %v, Исполнитель: %v, Описание: %v\n",
			j+1,
			task.Title,
			task.CreatorLogin,
			task.Author,
			task.Description,
		)
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
