package input

import (
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"path/filepath"
)

//TestRun makes validation tests on the run parameter
func TestRun(suite string) {
	if suite != "" {
		if !helper.DoesFileExist(suite) {
			Exit(fmt.Sprintf("\nrun parameter folder [%s] does not exist\n", suite))
		}
	} else {
		Exit("\nrun param is empty")
	}
}

//TestUI validates if the context folder exists
func TestUI(contextFolder string) {
	if contextFolder != "" {
		if !helper.DoesFileExist(contextFolder) {
			Exit(fmt.Sprintf("\nui parameter folder [%s] does not exist\n", contextFolder))
		}
	} else {
		Exit("\nui param is empty")
	}
}

//TestJavaHome returns a string if the javahome is valid, an empty if not
func TestJavaHome(home string) string {
	javaExecutable := "java"
	if home == "" {
		return javaExecutable
	}
	if helper.IsRunningOnWindows() {
		javaExecutable = filepath.Join(home, "bin", "java.exe")
	} else if helper.IsRunningOnLinux() {
		javaExecutable = filepath.Join(home, "bin", "java")
	} else {
		panic("Can not detect operatingsystem. Supported are: Windows, Linux")
	}

	if helper.DoesFileExist(home) {
		return javaExecutable
	}
	return ""
}
