package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var (
	mu       sync.Mutex
	filePath = "tasks.json"
)

func readTasks() ([]Task, error) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil //Если нет файла - возвращаем пустой список
		}
		return nil, fmt.Errorf("Не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("Не удалось прочитать задачи: %v", err)
	}

	return tasks, nil
}

func writeTasks(tasks []Task) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Не удалось создать файл: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("Не удалось записать задачи: %v", err)
	}

	return nil
}

func AddTask(text string, priority int) {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Ошибка при чтении задач")
		return
	}

	nextID := 1

	if len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}

	task := Task{
		ID:       nextID,
		Text:     text,
		Priority: priority,
	}

	tasks = append(tasks, task)

	if err := writeTasks(tasks); err != nil {
		fmt.Println("Ошибка при добавлении задачи: ", err)
	} else {
		fmt.Println("Задача добавлена успешно!")

	}
}

func ListTasks() ([]Task, error) {
	return readTasks()
}

func StatusSwitch(id int) bool {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Ошибка при чтении задач: ", err)
		return false
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = !tasks[i].Done
			if err := writeTasks(tasks); err != nil {
				fmt.Println("Ошибка при обновлении задачи: ", err)
				return false
			}
			return true
		}
	}

	return false
}

func EditTask(id int, text string, priority int) bool {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Ошибка при чтении задач: ", err)
		return false
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Text = text
			tasks[i].Priority = priority
			if err := writeTasks(tasks); err != nil {
				fmt.Println("Ошибка при обновлении задачи: ", err)
				return false
			}
			return true
		}
	}

	return false
}

func DeleteTask(id int) bool {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Ошибка при чтении задачи: ", err)
		return false
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := writeTasks(tasks); err != nil {
				fmt.Println("Ошибка при удалении задачи: ", err)
				return false
			}
			return true
		}
	}
	return false
}
