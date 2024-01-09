package fang

import (
	"github.com/atakanaydinbas/gofangdefang"
)

func DefangerCLI(input string) string {
	return gofangdefang.DefangAll(input)
}
