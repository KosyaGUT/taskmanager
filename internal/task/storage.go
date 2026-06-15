package task

import (
	"encoding/json"
	"os"
)

func SaveTasks(tasks []Task) error {
	file, err := os.Create("task.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}

func SaveUsers(users []User) error {
	file, err := os.Create("user.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(users)
	if err != nil {
		return err
	}
	return nil
}

func LoadTasks() ([]Task, error) {
	// Открываем файл для чтения
	file, err := os.Open("task.json")

	if os.IsNotExist(err) {
		return []Task{}, nil
	}

	if err != nil {
		return nil, err // Если файла нет или ошибка доступа, возвращаем её
	}
	defer file.Close()

	// Создаем переменную, куда запишем данные
	var tasks []Task

	// Создаем декодер и читаем данные из файла прямо в переменную
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks) // Обязательно передаем указатель (&)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func LoadUsers() ([]User, error) {
	// Открываем файл для чтения
	file, err := os.Open("user.json")

	if os.IsNotExist(err) {
		return []User{}, nil
	}

	if err != nil {
		return nil, err // Если файла нет или ошибка доступа, возвращаем её
	}
	defer file.Close()

	// Создаем переменную, куда запишем данные
	var users []User

	// Создаем декодер и читаем данные из файла прямо в переменную
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users) // Обязательно передаем указатель (&)
	if err != nil {
		return nil, err
	}

	return users, nil
}
