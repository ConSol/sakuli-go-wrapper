package input

import (
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"io/ioutil"
	"os"
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

//PrintVersion prints the sakuli version and env variables.
func PrintVersion() {
	versionFile := filepath.Join(helper.GetSahiHome(), "bin", "resources", "version.txt")
	data, err := ioutil.ReadFile(versionFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Print("\n--- Environment variables ---")
	fmt.Printf(`
SAKULI_HOME:                  %s
MOZ_DISABLE_OOP_PLUGINS:      %s
MOZ_DISABLE_AUTO_SAFE_MODE:   %s
MOZ_DISABLE_SAFE_MODE_KEY:    %s
`, helper.GetSahiHome(), os.Getenv("MOZ_DISABLE_OOP_PLUGINS"), os.Getenv("MOZ_DISABLE_AUTO_SAFE_MODE"), os.Getenv("MOZ_DISABLE_SAFE_MODE_KEY"))
	os.Exit(0)
}

func PrintExampleUsage() {
	fmt.Print(`Sakuli CLI Examples:

Usage: sakuli[.exe] COMMAND ARGUMENT [OPTIONS]

Run a test suite:
▶ Run the test suite "example":
    sakuli run "<your-project-path>/example"
▶ Use an infinite loop with 10 seconds pause between:
    sakuli run "<your-project-path>/example" -loop=10
▶ Choose browser "chrome" (browser must be registered):
    sakuli run "<your-project-path>/example" -browser=chrome
▶ Kill hanging processes in Windoes before:
    sakuli.exe run "<your-project-path>\example" -preHook='cscript.exe SAKULI_HOME\bin\helper\killproc.vbs -f SAKULI_HOME\bin\helper\procs_to_kill.txt'
▶ Run "exmaple_windows", increase the logging level:
    sakuli.exe run "<your-project-path>\example_windows" -D log.level.sakuli=DEBUG

Encrypt secrets:
▶ Encrypt a secret using an automatic determend NIC as salt:
    sakuli encrypt topsecret -interface auto
▶ Encrypt a secret using eth0 as salt NIC:
    sakuli encrypt topsecret -interface eth0
▶ Show interfaces available for encryption:
    sakuli encrypt topsecret -interface list

Others:
▶ Show version (use this information when submitting bugs):
    sakuli -version
`)
	os.Exit(0)
}