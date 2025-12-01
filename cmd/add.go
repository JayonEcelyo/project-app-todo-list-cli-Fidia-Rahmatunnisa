package cmd

import (
	"fmt"
	"mini-project-2/service"
	"github.com/spf13/cobra"
)

var title, status, priority string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Tambah tugas baru",
	Run: func(cmd *cobra.Command, args []string) {
		if title == "" {
			fmt.Println("task tidak boleh kosong")
			return
		}

		err := service.AddTask(title, status, priority)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Tugas berhasil ditambahkan!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title, "task", "t", "", "Judul tugas")
	addCmd.Flags().StringVarP(&status, "status", "s", "new", "Status tugas(new/progress/pending/completed)")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "low", "Prioritas tugas (low/medium/high)")
}
