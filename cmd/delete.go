package cmd

import (
	"fmt"
	"mini-project-2/service"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Hapus tugas",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		err := service.DeleteTask(id)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task dihapus")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
