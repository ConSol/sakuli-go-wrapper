package execute

import (
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"os"
	"runtime"
)

//RunHandler runs external program with no parameters
func RunHandler(executable string) {
	fmt.Printf("Calling Handler executable: %s\n", executable)
	var returnCode int
	var err error
	if helper.IsRunningOnWindows() {
		returnCode, err = Execute("cmd", "/c", executable)
	} else if helper.IsRunningOnLinux() {
		returnCode, err = Execute("sh", "-c", executable)
	} else {
		panic("Unkown OS: " + runtime.GOOS)
	}
	if err != nil {
		fmt.Printf("Error while calling: %s\n", executable)
	}
	fmt.Printf("Handler [%s] finished with returncode: %d\n", executable, returnCode)
	if returnCode != 0 {
		os.Exit(returnCode)
	}
}
