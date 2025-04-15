package cmd

import (
	"bufio"
	//"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var toDO string
var listed task

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(toDO)

		title := toDO + ".txt"

		_, err := os.Stat(title)
		if os.IsNotExist(err) {
			fmt.Printf("To-Do list %v does not exist: \n", toDO)
			return
		}
		
		file, err := os.Open(title)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			data := []byte(scanner.Text())
			// err := json.Unmarshal(data, &listed)
			// if err != nil {
			// 	fmt.Printf("Error unmarshalling JSON: %v\n", err)
			// 	continue
			// }
			fmt.Println(string(data))
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}

	},
}

func init() {
	taskCmd.AddCommand(taskListCmd)

	taskListCmd.Flags().StringVarP(&toDO, "to-do", "t", "", "To-Do list name (required)")
	taskListCmd.MarkFlagRequired("to-do")
}
