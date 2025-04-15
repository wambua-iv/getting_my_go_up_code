package cmd


import (
	"fmt"
	"os"
	"time"
	"encoding/json"
	
	"github.com/spf13/cobra"
)


func init() {
	taskCmd.AddCommand(taskAddCmd)
	taskAddCmd.Flags().StringVarP(&toDo, "to-do", "t", "", "To-Do list name (required)")
	taskAddCmd.Flags().StringVarP(&name, "name", "n", "", "Task name (required)")
	taskAddCmd.MarkFlagRequired("to-do")
	taskAddCmd.MarkFlagRequired("name")
}

var taskAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  "Add a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		// if len(args) == 0 {
		// 	fmt.Println("Please provide a task description")
		// 	os.Exit(1)
		// }
		// taskName := args[0]
		fmt.Printf("Added task: %s\n", name)
		file := toDo + ".txt"
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			fmt.Printf("To Do list: '%v' not created \n", toDO)
			return
		}
		f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0643)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()
		
		tasks:= task{
			Id: "looping",
			Name: name,
			Status: "pending",
			Description: "",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		
		encode := json.NewEncoder(f)
		encode.SetIndent(""," ")
		if err := encode.Encode(tasks); err != nil {
			fmt.Println("Error encoding task:", err)
			return
		}

		if verbose {
			fmt.Println("Task created successfully")
		}
	},
}