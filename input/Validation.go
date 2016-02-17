package input

import (
	"flag"
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"os"
	"path/filepath"
)

//TestRun makes validation tests on the run parameter
func TestRun(suite string) {
	if suite == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "\nrun param is empty")
		os.Exit(999)
	} else if !helper.DoesFileExist(suite) {
		fmt.Fprintf(os.Stderr, "run parameter folder [%s] does not exitst\n", suite)
		os.Exit(999)
	}
}

//TestJavaHome returns a string if the javahome is valid, an empty if not
func TestJavaHome(home string) string {
	javaExecutable := "java"
	if home == "" {
		return javaExecutable
	}
	if helper.IsRunningOnWindows() {
		javaExecutable = filepath.Join(home, "java.exe")
	} else if helper.IsRunningOnLinux() {
		javaExecutable = filepath.Join(home, "java")
	} else {
		panic("Can not detect operatingsystem. Supported are: Windows, Linux")
	}

	if helper.DoesFileExist(home) {
		return javaExecutable
	}
	return ""
}
