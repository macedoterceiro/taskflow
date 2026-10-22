package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func main() {
	addCmd := flag.String("add", "", "Adiciona uma nova tarefa")
	listCmd := flag.Bool("list", false, "Lista todas as tarefas")
	flag.Parse()

	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Erro ao carregar tarefas: %v\n", err)
		os.Exit(1)
	}

	if *addCmd != "" {
		newTask := Task{
			ID:        len(tasks) + 1,
			Title:     *addCmd,
			Completed: false,
		}
		tasks = append(tasks, newTask)
		if err := saveTasks(tasks); err != nil {
			fmt.Printf("Erro ao salvar tarefa: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Tarefa adicionada com sucesso!")
		return
	}

	if *listCmd {
		if len(tasks) == 0 {
			fmt.Println("Nenhuma tarefa encontrada.")
			return
		}
		fmt.Println("\n--- SUAS TAREFAS ---")
		for _, t := range tasks {
			status := "[ ]"
			if t.Completed {
				status = "[X]"
			}
			fmt.Printf("%d. %s %s\n", t.ID, status, t.Title)
		}
		return
	}

	fmt.Println("Uso: go run main.go -add \"Minha tarefa\" ou go run main.go -list")
}
