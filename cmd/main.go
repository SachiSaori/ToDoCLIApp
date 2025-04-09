package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"ToDoListCLIApp/internal/todo"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("--Меню--")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу выполненной")
		fmt.Println("4. Удалить задачу")
		fmt.Println("0. Выйти")
		fmt.Println("> ")

		reader.Scan()
		choice := reader.Text()

		switch choice {
		case "1":
			fmt.Println("Текст задачи: ")
			reader.Scan()
			text := reader.Text()

			fmt.Println("Приоритет числом (1 - наивысший): ")
			reader.Scan()
			priority, _ := strconv.Atoi(reader.Text())

			todo.AddTask(text, priority)
		case "2":
			todo.ListTasks()
		case "3":
			fmt.Println("Введите ID задачи которую хотите пометить как выполненную: ")
			reader.Scan()
			id, _ := strconv.Atoi(reader.Text())
			if todo.MarkDone(id) {
				fmt.Println("Задача отмечена как выполненная!")
			} else {
				fmt.Println("Задача не найдена!")
			}
		case "4":
			fmt.Println("Введите ID задачи которую хотите удалить: ")
			reader.Scan()
			id, _ := strconv.Atoi(reader.Text())
			if todo.DeleteTask(id) {
				fmt.Println("Задача успешно удалена!")
			} else {
				fmt.Println("Задача не найдена!")
			}
		case "0":
			fmt.Println("До встречи")
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
	}

}
