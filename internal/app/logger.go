package app

import (
	"fmt"
)

//Debug print trace debug
func Debug(str string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("@debug: %s\n", str), a...)
}

//Warning print trace warning
func Warning(str string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("@warning: %s\n", str), a...)
}

//Error print trace error
func Error(str string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("@error: %s\n", str), a...)
}

//RunScript print trace error
func RunScript(str string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("@run: %s\n", str), a...)
}
