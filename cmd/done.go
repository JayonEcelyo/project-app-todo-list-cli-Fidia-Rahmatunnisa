package cmd

import (
	"fmt"
	"mini-project-2/service"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Tandai tugas selesai",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Masukkan ID tugas")
			return
		}

		id, _ := strconv.Atoi(args[0])
		err := service.MarkDone(id)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Tugas selesai!")
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
