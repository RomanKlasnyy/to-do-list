package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	ID       uint `gorm:"primaryKey"`
	Task     string
	Deadline time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Task{})

	for {
		var choice string
		fmt.Println("1) Today's tasks")
		fmt.Println("2) Week's tasks")
		fmt.Println("3) All tasks")
		fmt.Println("4) Missed tasks")
		fmt.Println("5) Add task")
		fmt.Println("6) Delete task")
		fmt.Println("0) Exit")
		fmt.Scan(&choice)

		switch choice {
		case "0":
			os.Exit(0)
		case "1":
			today := time.Now()
			rows := []Task{}
			db.Where("deadline BETWEEN ? AND ?", today, today.Add(24*time.Hour)).Find(&rows)
			printTasks(rows)
		case "2":
			for i := 0; i < 7; i++ {
				day := time.Now().AddDate(0, 0, i)
				rows := []Task{}
				db.Where("deadline BETWEEN ? AND ?", day, day.Add(24*time.Hour)).Find(&rows)
				fmt.Printf("%s %s %d:\n", day.Weekday(), day.Month(), day.Day())
				printTasks(rows)
				fmt.Println()
			}
		case "3":
			rows := []Task{}
			db.Order("deadline").Find(&rows)
			printTasksWithDate(rows)
		case "4":
			today := time.Now()
			rows := []Task{}
			db.Where("deadline < ?", today).Find(&rows)
			printTasksWithDate(rows)
		case "5":
			addTask(db)
		case "6":
			deleteTask(db)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Nothing to do!")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Task)
	}
}

func printTasksWithDate(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Nothing to do!")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s. %s\n", i+1, task.Task, task.Deadline.Format("02 Jan"))
	}
}

func addTask(db *gorm.DB) {
	var taskString, deadlineString string
	fmt.Println("Enter task:")
	fmt.Scan(&taskString)
	fmt.Println("Enter deadline (YYYY-MM-DD):")
	fmt.Scan(&deadlineString)

	deadline, err := time.Parse("2006-01-02", deadlineString)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD format.")
		return
	}

	task := Task{
		Task:     taskString,
		Deadline: deadline,
	}
	db.Create(&task)
	fmt.Println("The task has been added!")
}

func deleteTask(db *gorm.DB) {
	var idString string
	fmt.Println("Choose the number of the task you want to delete:")
	fmt.Scan(&idString)
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	var task Task
	db.First(&task, id)
	db.Delete(&task)
	fmt.Println("The task has been deleted!")
}
