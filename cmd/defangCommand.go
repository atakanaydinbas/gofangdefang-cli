package gofangdefang

import (
	"fmt"

	defang "github.com/atakanaydinbas/gofangdefang-cli/pkg/defang"
	"github.com/spf13/cobra"
)

var defangCMD = &cobra.Command{
	Use:     "defang",
	Aliases: []string{"def", "d"},
	Short:   "Defang an IOC",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(defang.DefangerCLI(args[0]))
	},
}
var DefangFileCMD = &cobra.Command{
	Use:   "defangfile -i <filepath> -o <outputpath>",
	Short: "Defang an IOC from a file path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(DefangFileInput)
		fmt.Println(DefangFileoutput)
		fmt.Println(defang.DefangFileCLI(DefangFileInput, DefangFileoutput))
	},
}
var DefangFileInput string
var DefangFileoutput string

func init() {
	DefangFileCMD.Flags().StringVarP(&DefangFileInput, "input", "i", "", "Input file path to defang")
	DefangFileCMD.MarkFlagRequired("input")
	DefangFileCMD.MarkFlagFilename("input")
	DefangFileCMD.Flags().StringVarP(&DefangFileoutput, "output", "o", "", "Output file path to defang")
	DefangFileCMD.MarkFlagRequired("output")
	DefangFileCMD.MarkFlagFilename("output")
	rootCmd.AddCommand(DefangFileCMD)
	rootCmd.AddCommand(defangCMD)
}
