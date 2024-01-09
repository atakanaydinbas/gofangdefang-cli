package gofangdefang

import (
	"fmt"

	defang "github.com/atakanaydinbas/gofangdefang-cli/pkg/defang"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "defang",
	Aliases: []string{"def", "d"},
	Short:   "Defang an IOC",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(defang.DefangerCLI(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
