package gofangdefang

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofangdefang-cli",
	Short: "GoFangDefang is a tool for defanging and fanging IOCs",
	Long:  `GoFangDefang is a CLI tool for secure manipulation of Indicators of Compromise (IOCs), converting them between their original "fang" format (with special characters) and a safer "defang" format. It prevents accidental execution of potentially malicious IOCs like URLs, IPs, domains, or subdomains.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
