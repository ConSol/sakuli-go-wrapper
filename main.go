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
	var inter string
	var sahiHome string
	var version bool

	sakuliJars := filepath.Join(helper.GetSahiHome(), "libs", "java")
	myFlagSet := flag.NewFlagSet("", flag.ExitOnError)
	input.MyFlagSet = myFlagSet
	myFlagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, `Generic Sakuli test starter.
%d - The Sakuli team <sakuli@consol.de>
http://www.sakuli.org
https://github.com/ConSol/sakuli

Usage: sakuli[.exe] COMMAND ARGUMENT [OPTIONS]

       sakuli -help
       sakuli -version
       sakuli run <sakuli suite path> [OPTIONS]
       sakuli encrypt <secret> [OPTIONS]

Commands:
       run 	   <sakuli suite path>
       encrypt 	   <secret>

Options:
       -loop	   <seconds>	  Loop this suite, wait n seconds between
                                  executions, 0 means no loops (default: 0)
       -javaHome   <folder>       Java bin dir (overrides PATH)
       -javaOption <java option>  JVM option parameter, e.g. '-agentlib:...'
       -preHook    <programpath>  A program which will be executed before a
                                  suite run (can be added multiple times)
       -postHook   <programpath>  A program which will be executed after a
                                  suite run (can be added multiple times)
       -D 	   <JVM option>   JVM option to set a property at runtime,
                                  overrides file based properties
       -browser    <browser>      Browser for the test execution
                                  (default: Firefox)
       -interface  <interface>    Network interface card name, used by
                                  command 'encrypt' as salt
       -sahiHome   <folder>       Sahi installation folder
       -version                   Version info
       -help                      This help text

Examples: 
    * Run the test suite "example": 
    sakuli run "$SAKULI_HOME\..\suites\example"
    * Run "example" in an infinite loop with 10 seconds pause between: 
    sakuli run "$SAKULI_HOME\..\suites\example" -loop=10
    * Run "example" with browser "chrome" (browser must be registered): 
    sakuli run "$SAKULI_HOME\..\suites\example" -browser=chrome
    * Run "example", kill hanging processes before:
    sakuli run "$SAKULI_HOME\..\suites\example" \
      -preHook='cscript.exe $SAKULI_HOME\bin\helper\killproc.vbs     \
      -f $SAKULI_HOME\bin\helper\procs_to_kill.txt'
    * Run "exmaple_windows", increase the logging level: 
    sakuli run "$SAKULI_HOME\..\suites\example" -D log.level.sakuli=DEBUG

    * Encrypt a secret using eth0 as salt NIC: 
    sakuli encrypt topsecret -interface eth0
    * Show interfaces available for encryption: 
    sakuli encrypt topsecret -interface list

    * Show version (use this information when submitting bugs): 
    sakuli -version

`, time.Now().Year())
	}

	myFlagSet.IntVar(&loop, "loop", 0, "loop this suite, wait n seconds between executions, 0 means no loops (default: 0)")
	myFlagSet.StringVar(&javaHome, "javaHome", "", "Java bin dir (overrides PATH)")
	myFlagSet.Var(&preHooks, "preHook", "A program which will be executed before a suite run (can be added multiple times)")
	myFlagSet.Var(&postHooks, "postHook", "A program which will be executed after a suite run (can be added multiple times)")

	myFlagSet.Var(&javaProperties, "D", "JVM option to set a property at runtime, overrides file based properties")
	myFlagSet.Var(&javaOptions, "javaOption", "JVM option parameter, e.g. '-agentlib:...'")
	myFlagSet.StringVar(&browser, "browser", "", "browser for the test execution (default: Firefox)")
	myFlagSet.StringVar(&inter, "interface", "", "network interface icaed name, used by command 'encrypt' as salt")
	myFlagSet.StringVar(&sahiHome, "sahi_home", "", "Sahi installation folder")
	myFlagSet.BoolVar(&version, "version", false, "version info")

	if len(os.Args) > 2 {
		myFlagSet.Parse(os.Args[3:])
	} else {
		myFlagSet.Parse(os.Args[1:])
		if version {
			input.PrintVersion()
		}
		input.ExitWithHelp("\nOnly 'sakuli COMMAND ARGUMENT [OPTIONS]' is allowed, given: " + fmt.Sprint(os.Args))
	}

	sakuliProperties := map[string]string{"sakuli_home": helper.GetSahiHome()}
	typ, argument := input.ParseArgs(append(os.Args[1:3],myFlagSet.Args()...))
	switch typ {
	case input.RunMode:
		input.TestRun(argument)
		sakuliProperties[input.RunMode] = argument
	case input.EncryptMode:
		sakuliProperties[input.EncryptMode] = argument
	case input.Error:
		panic("can't pars args")
	}

	javaExecutable := input.TestJavaHome(javaHome)
	javaProperties = javaProperties.AddPrefix("-D")

	if browser != "" {
		sakuliProperties["browser"] = browser
	}
	if inter != "" {
		sakuliProperties["interface"] = inter
	}
	if sahiHome != "" {
		sakuliProperties["sahiHome"] = sahiHome
	}
	joinedSakuliProperties := genSakuliPropertiesList(sakuliProperties)

	if len(preHooks) > 0 {
		fmt.Println("=========== Starting Pre-Hooks ===========")
		for _, pre := range preHooks {
			execute.RunHandler(pre)
		}
		fmt.Println("=========== Finished Pre-Hooks ===========")
	}

	sakuliReturnCode := execute.RunSakuli(javaExecutable, sakuliJars, javaOptions, javaProperties, joinedSakuliProperties)
	for loop > 0 {
		fmt.Printf("*** Loop mode - sleeping for %d seconds... ***\n", loop)
		time.Sleep(time.Duration(loop) * time.Second)
		execute.RunSakuli(javaExecutable, sakuliJars, javaOptions, javaProperties, joinedSakuliProperties)
	}

	if len(postHooks) > 0 {
		fmt.Println("=========== Starting Post-Hooks ===========")
		for _, post := range postHooks {
			execute.RunHandler(post)
		}
		fmt.Println("=========== Finished Post-Hooks ===========")
	}
	os.Exit(sakuliReturnCode)
}

func genSakuliPropertiesList(properties map[string]string) input.StringSlice {
	propertiesString := []string{}
	for k, v := range properties {
		propertiesString = append(propertiesString, fmt.Sprintf("--%s", k))
		propertiesString = append(propertiesString, v)
	}
	return propertiesString
}
