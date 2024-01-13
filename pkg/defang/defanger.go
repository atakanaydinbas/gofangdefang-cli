package fang

import (
	"github.com/atakanaydinbas/gofangdefang"
)

func DefangerCLI(input string) string {
	return gofangdefang.DefangAll(input)
}

func DefangFileCLI(inputpath string, outputpath string) string {
	if inputpath == "" || outputpath == "" {
		return "Input and output paths cannot be empty"
	}
	_, err := gofangdefang.DefangFile(inputpath, true, outputpath)
	if err != nil {
		return err.Error()
	}
	return "File fanged successfully"
}
