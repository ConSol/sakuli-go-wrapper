package input

import (
	"flag"
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"os"
)

const (
	//RunMode is used for executing a test
	RunMode = "run"
	//UiMode is used to start sakuli UI jar
	UiMode = "ui"
	//EncryptMode is used to encrypt a password
	EncryptMode = "encrypt"
	//CreateMode is used for creating new objects like a masterkey
	CreateMode = "create"
	//Error this should never happen
	Error = "should never happen"

	//OPTION values
	//key for the sakuli home value
	OptSakuliHome = "sakuliHome"
	OptSahiHome   = "sahiHome"
	OptBrowser    = "browser"
	OptInterface  = "interface"
	OptMasterkey  = "masterkey"
)

//ParseArgs parses the COMMAND and its options
func ParseArgs(args []string) (string, string) {
	length := len(args)

	containsRun, indexRun := helper.Contains(args, RunMode)
	containsUi, indexUi := helper.Contains(args, UiMode)
	containsEncrypt, indexEncrypt := helper.Contains(args, EncryptMode)
	containsCreate, indexCreate := helper.Contains(args, CreateMode)

	detError := ""
	if !containsRun && !containsUi && !containsEncrypt && !containsCreate {
		detError += "Incorrect COMMAND, please use a valid one: "
	}
	if len(detError) > 0 || length != 2 {
		ExitWithHelp("\n" + detError + "Only 'sakuli COMMAND ARGUMENT [OPTIONS]' is allowed, given: " + fmt.Sprint(args))
	}

	if containsRun {
		return RunMode, args[(indexRun+1)%length]
	} else if containsUi {
		return UiMode, args[(indexUi+1)%length]
	} else if containsEncrypt {
		return EncryptMode, args[(indexEncrypt+1)%length]
	} else if containsCreate {
		return CreateMode, args[(indexCreate+1)%length]
	}
	return Error, Error
}

//MyFlagSet contains the parsed arguemnts, will be set at the beginning.
var MyFlagSet *flag.FlagSet

//ExitWithHelp prints the help, the info and exits the program with 999.
func ExitWithHelp(info string) {
	MyFlagSet.Usage()
	fmt.Fprintln(os.Stdout, "================================================================================")
	fmt.Fprintln(os.Stderr, info)
	os.Exit(999)
}

//ExitWithHelp prints the help, the info and exits the program with 999.
func Exit(info string) {
	fmt.Fprintln(os.Stderr, info)
	os.Exit(999)
}
