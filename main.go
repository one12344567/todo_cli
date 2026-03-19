package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const FILENAME = "todos.json"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入具体参数")
		return
	}
	switch os.Args[1] {
	case "add":
		fmt.Println("添加todo")
		addtodo(os.Args[2], os.Args[3])
	case "list":
		fmt.Println("展示todo list")
		fmt.Println(listtodos())
	case "delete":
		fmt.Println("删除事项")
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id 必须是数字")
			return
		}

		deletetodo(id)
	}
}

func addtodo(title string, description string) {
	todos := loadtodos()
	new_todo := Todo{
		ID:          nextID(todos),
		Title:       title,
		Description: description,
		Completed:   false,
	}
	todos = append(todos, new_todo)
	savetodo(todos)
}

func savetodo(todos []Todo) {
	data, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("savetodo err!", err)
	}
	err2 := os.WriteFile(FILENAME, data, 0644)
	if err2 != nil {
		fmt.Println("writefile err!", err2)
	}
}

func listtodos() string {
	jsondata, err := os.ReadFile(FILENAME)
	if err != nil {
		fmt.Println("读取json错误")
	}
	var godata []Todo
	err2 := json.Unmarshal(jsondata, &godata)
	if err2 != nil {
		fmt.Println("json->go转换错误", err2)
	}
	result := string(jsondata)
	return result
}

func loadtodos() []Todo {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		fmt.Println("load err!", err)
	}
	if len(data) == 0 {
		return []Todo{}
	}
	var todos []Todo
	err2 := json.Unmarshal(data, &todos)
	if err2 != nil {
		fmt.Println("loadtodos err2!", err2)
		return []Todo{}
	}
	return todos
}

func deletetodo(id int) {

	todos := loadtodos()
	for index, item := range todos {
		if item.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			savetodo(todos)
			return
		}
	}
	fmt.Println("delete err")
}

func nextID(todos []Todo) int {
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID + 1
}
