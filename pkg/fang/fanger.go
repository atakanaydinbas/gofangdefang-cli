package fang

import (
	"github.com/atakanaydinbas/gofangdefang"
)

func FangerCLI(input string) string {
	return gofangdefang.FangAll(input)
}

func FangFileCLI(inputpath string, outputpath string) string {
	if inputpath == "" || outputpath == "" {
		return "Input and output paths cannot be empty"
	}
	_, err := gofangdefang.FangFile(inputpath, true, outputpath)
	if err != nil {
		return err.Error()
	}
	return "File fanged successfully"
}
