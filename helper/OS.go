package helper

import (
	"fmt"
	"github.com/kardianos/osext"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//IsRunningOnWindows returns true if the program is running on Windows
func IsRunningOnWindows() bool {
	return runtime.GOOS == "windows"
}

//IsRunningOnLinux returns true if the program is running on Linux
func IsRunningOnLinux() bool {
	return runtime.GOOS == "linux"
}

//GenClassPath concatenates the elements with ; on Windows with : on Linux
func GenClassPath(elements ...string) string {
	if IsRunningOnWindows() {
		return strings.Join(elements, ";")
	} else if IsRunningOnLinux() {
		return strings.Join(elements, ":")
	} else {
		panic("Can not detect operatingsystem. Supported are: Windows, Linux")
	}
}

//DoesFileExist returns true if the given file exists
func DoesFileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

var sakuliHome = ""

//GetSakuliHome returns the sakulihome folder.
//First lookup: env: SAKULI_HOME
//Second: one folder above the binary
func GetSakuliHome() string {
	if sakuliHome == "" {
		sakuliHome = os.Getenv("SAKULI_HOME")
		if sakuliHome == "" {
			var err error
			var execFolder string
			execFolder, err = osext.ExecutableFolder()
			sakuliHome, err = filepath.Abs(filepath.Join(execFolder, ".."))
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(os.Stderr, "=================== SAKULI_HOME is empty using binary folder ===================\n"+sakuliHome+"\n================================================================================")
		}
	}
	return sakuliHome
}

//GetSakuliRoot returns the sakuliroot folder.
//lookup: env: SAKULI_ROOT
func GetSakuliRoot() string {
	var sakuliRoot = os.Getenv("SAKULI_ROOT")
	if sakuliHome == "" {
		panic("environment variable 'SAKULI_ROOT' is not defined!")
	}
	return sakuliRoot
}
