/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"
	uuid "github.com/google/uuid"
	"github.com/spf13/cobra"
)
var toDo string
var name string
var verbose bool

type task struct {
	id string
	name string
	status string
	description string
	createdAt int64
	updatedAt int64
}

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(toDo)
		file := toDo + ".txt"
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			fmt.Println("To Do list not created")
			return
		}
		f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()
		
		task := task{
			id: uuid.New().String(),
			name: name,
			status: "pending",
			description: "",
			createdAt: time.Now().Unix(),
			updatedAt: time.Now().Unix(),
		}

		if _, err := f.WriteString(fmt.Sprintf("%v\n", task)); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		if verbose {
			fmt.Println("Task created successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
	
	taskCmd.Flags().StringVarP(&toDo, "to-do", "t", "", "To-Do list name (required)")
	taskCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the task is required (required)")
	taskCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	
	taskCmd.MarkFlagRequired("to-do")
	taskCmd.MarkFlagRequired("name")
	

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
