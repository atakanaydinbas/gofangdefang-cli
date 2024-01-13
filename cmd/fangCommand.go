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

var fangFileCMD = &cobra.Command{
	Use:   "fangfile -i <inputfilepath> -o <outputfilepath>",
	Short: "Fang an IOC from a file path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(FangFileinput)
		fmt.Println(FangFileoutput)
		fmt.Println(fang.FangFileCLI(FangFileinput, FangFileoutput))
	},
}
var FangFileinput string
var FangFileoutput string

func init() {
	fangFileCMD.Flags().StringVarP(&FangFileinput, "input", "i", "", "Input file path to fang")
	fangFileCMD.MarkFlagRequired("input")
	fangFileCMD.MarkFlagFilename("input")
	fangFileCMD.Flags().StringVarP(&FangFileoutput, "output", "o", "", "Output file path to fang")
	fangFileCMD.MarkFlagRequired("output")
	fangFileCMD.MarkFlagFilename("output")

	rootCmd.AddCommand(fangCMD)
	rootCmd.AddCommand(fangFileCMD)
}
