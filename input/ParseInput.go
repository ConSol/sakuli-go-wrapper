package input

import (
	"flag"
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"os"
)

const (
	//EncryptMode is used to encrypt a password
	EncryptMode = "encrypt"
	//RunMode is used for executing a test
	RunMode = "run"
	//Error this should never happen
	Error = "should never happen"
)

//ParseArgs parses the COMMAND and its options
func ParseArgs(args []string) (string, string) {
	length := len(args)
	if length != 2 {
		ExitWithHelp("\nOnly COMMAND + ARGUMENT + OPTIONS is allowed given: " + fmt.Sprint(args))
	}
	containsEncrypt, indexEncrypt := helper.Contains(args, EncryptMode)
	containsRun, indexRun := helper.Contains(args, RunMode)

	if !containsEncrypt && !containsRun {
		ExitWithHelp("\nrun and encrypt are missing")
	}
	if containsEncrypt && containsRun {
		ExitWithHelp("\nrun and encrypt are given, only one is needed")
	}

	if containsEncrypt {
		return EncryptMode, args[(indexEncrypt+1)%length]
	} else if containsRun {
		return RunMode, args[(indexRun+1)%length]
	}
	return Error, Error
}

//ExitWithHelp prints the help, the info and exits the program with 999.
func ExitWithHelp(info string) {
	flag.Usage()
	fmt.Fprintln(os.Stderr, info)
	os.Exit(999)
}
