package main

import (
	"flag"
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/execute"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"github.com/ConSol/sakuli-go-wrapper/input"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var loop int
	var javaHome string
	var javaProperties input.StringSlice
	var javaOptions input.StringSlice
	var preHooks input.StringSlice
	var postHooks input.StringSlice
	var browser string
	var sahiHome string
	var inter string
	var masterkey string
	var version bool
	var examples bool

	sakuliJars := filepath.Join(helper.GetSakuliHome(), "libs", "java")
	sakuliUiJar := filepath.Join(helper.GetSakuliHome(), "libs", "ui", "java", "sakuli-ui-web.jar")
	uiInstalled := helper.DoesFileExist(sakuliUiJar)

	myFlagSet := flag.NewFlagSet("", flag.ExitOnError)
	input.MyFlagSet = myFlagSet
	myFlagSet.Usage = func() {
		input.PrintHelp(uiInstalled)
	}

	myFlagSet.IntVar(&loop, "loop", 0, "loop this suite, wait n seconds between executions, 0 means no loops (default: 0)")
	myFlagSet.StringVar(&javaHome, "javaHome", "", "Java bin dir (overwrites PATH)")
	myFlagSet.Var(&preHooks, "preHook", "A program which will be executed before a suite run (can be added multiple times)")
	myFlagSet.Var(&postHooks, "postHook", "A program which will be executed after a suite run (can be added multiple times)")

	myFlagSet.Var(&javaProperties, "D", "JVM option to set a property at runtime, overwrites file based properties")
	myFlagSet.Var(&javaOptions, "javaOption", "JVM option parameter, e.g. '-agentlib:...'")
	myFlagSet.StringVar(&browser, "browser", "", "browser for the test execution (default: Firefox)")
	myFlagSet.StringVar(&masterkey, "masterkey", "", "AES base64 key used by command 'encrypt'")
	myFlagSet.StringVar(&inter, "interface", "", "network interface icaed name, used by command 'encrypt' as salt")
	myFlagSet.StringVar(&sahiHome, "sahiHome", "", "Sahi installation folder")
	myFlagSet.BoolVar(&examples, "examples", false, "CLI usage examples")
	myFlagSet.BoolVar(&version, "version", false, "version info")

	if len(os.Args) > 2 {
		myFlagSet.Parse(os.Args[3:])
	} else {
		myFlagSet.Parse(os.Args[1:])
		if version {
			input.PrintVersion()
		}
		if examples {
			input.PrintExampleUsage(uiInstalled)
		}
		detError := ""
		if len(os.Args) == 2 {
			detError += "ARGUMENT is missing specify one: "
		}
		input.ExitWithHelp("\n" + detError + "Only 'sakuli COMMAND ARGUMENT [OPTIONS]' is allowed, given: " + fmt.Sprint(os.Args))
	}

	sakuliProperties := map[string]string{input.OptSakuliHome: helper.GetSakuliHome()}
	typ, argument := input.ParseArgs(append(os.Args[1:3], myFlagSet.Args()...))
	switch typ {
	case input.RunMode:
		input.TestRun(argument)
		sakuliProperties[input.RunMode] = argument
	case input.UiMode:
		if !uiInstalled {
			input.ExitWithHelp("\nSakuli UI is NOT INSTALLED!\n" +
				"Only 'sakuli COMMAND ARGUMENT [OPTIONS]' is allowed, given: " + fmt.Sprint(os.Args))
		}
		input.TestUI(argument)
		sakuliProperties[input.UiMode] = argument
	case input.CreateMode:
		sakuliProperties[input.CreateMode] = argument
	case input.EncryptMode:
		sakuliProperties[input.EncryptMode] = argument
	case input.Error:
		panic("can't pars args")
	}

	javaExecutable := input.TestJavaHome(javaHome)
	javaProperties = javaProperties.AddPrefix("-D")

	if browser != "" {
		sakuliProperties[input.OptBrowser] = browser
	}
	if inter != "" {
		sakuliProperties[input.OptInterface] = inter
	}
	if masterkey != "" {
		sakuliProperties[input.OptMasterkey] = masterkey
	}
	if sahiHome != "" {
		sakuliProperties[input.OptSahiHome] = sahiHome
	}

	if len(preHooks) > 0 {
		fmt.Println("=========================== Starting Pre-Hooks =================================")
		for _, pre := range preHooks {
			execute.RunHandler(pre)
		}
		fmt.Println("=========================== Finished Pre-Hooks =================================")
	}

	sakuliReturnCode := 0
	if _, startUI := sakuliProperties[input.UiMode]; startUI {
		sakuliReturnCode = execute.RunSakuliUI(javaExecutable, sakuliUiJar, javaOptions, javaProperties, sakuliProperties)
	} else {
		sakuliReturnCode = execute.RunSakuli(javaExecutable, sakuliJars, javaOptions, javaProperties, sakuliProperties)
		for loop > 0 {
			fmt.Printf("*** Loop mode - sleeping for %d seconds... ***\n", loop)
			time.Sleep(time.Duration(loop) * time.Second)
			execute.RunSakuli(javaExecutable, sakuliJars, javaOptions, javaProperties, sakuliProperties)
		}
	}
	if len(postHooks) > 0 {
		fmt.Println("=========================== Starting Post-Hooks ================================")
		for _, post := range postHooks {
			execute.RunHandler(post)
		}
		fmt.Println("=========================== Finished Post-Hooks ================================")
	}
	os.Exit(sakuliReturnCode)
}
