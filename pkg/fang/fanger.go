package defang

import (
	"github.com/atakanaydinbas/gofangdefang"
)

func FangerCLI(input string) string {
	return gofangdefang.FangAll(input)
}
