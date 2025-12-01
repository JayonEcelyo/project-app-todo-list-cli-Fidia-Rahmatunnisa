package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mini-project-2/service"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Tampilkan daftar tugas",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := service.ListTasks()

		fmt.Println("================================ List Tasks ===============================")

		border := "+----+----------------------+------------+----------+"
		fmt.Println(border)
		fmt.Printf("| %-2s | %-20s | %-10s | %-6s |\n", "No", "Task", "Status", "Priority")
		fmt.Println(border)

		for _, t := range tasks {
			plainStatus := t.Status
			plainPriority := t.Priority

			status := plainStatus
			priority := plainPriority

			// Apply color after spacing rules
			switch plainStatus {
			case "completed":
				status = Blue + plainStatus + Reset
			case "progress":
				status = Yellow + plainStatus + Reset
			case "pending":
				status = Red + plainStatus + Reset
			default:
				status = Green + plainStatus + Reset
			}

			switch plainPriority {
			case "high":
				priority = Red + plainPriority + Reset
			case "medium":
				priority = Blue + plainPriority + Reset
			default:
				priority = Yellow + plainPriority + Reset
			}

			fmt.Printf(
				"| %-2d | %-20s | %-11s | %-8s |\n",
				t.ID, t.Task, status, priority,
			)
		}

		fmt.Println(border)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
