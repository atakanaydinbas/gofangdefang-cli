package gofangdefang

import (
	"fmt"

	fang "github.com/atakanaydinbas/gofangdefang-cli/pkg/fang"
	"github.com/spf13/cobra"
)

var fangCMD = &cobra.Command{
	Use:     "fang",
	Aliases: []string{"f"},
	Short:   "Fang an IOC",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fang.FangerCLI(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(fangCMD)
}
