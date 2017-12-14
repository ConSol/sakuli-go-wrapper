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
	myFlagSet := flag.NewFlagSet("", flag.ExitOnError)
	input.MyFlagSet = myFlagSet
	myFlagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, `Generic Sakuli test starter.
%d - The Sakuli team <sakuli@consol.de>
http://www.sakuli.org
https://github.com/ConSol/sakuli

Usage: sakuli[.exe] COMMAND ARGUMENT [OPTIONS]

  sakuli -help
  sakuli -examples
  sakuli -version
  sakuli run       <sakuli suite path> [OPTIONS]
  sakuli encrypt   <secret> [OPTIONS]
  sakuli create    <object> [OPTIONS]

Commands:
  run              <sakuli suite path>
  encrypt          <secret>
  create           <object>

Objects:
  masterkey        Base64 decoded AES 128-bit key (for encryption)

Options:
  -loop            <seconds>           Loop this suite, wait n seconds between
                                       executions, 0 means no loops (default: 0)
  -javaHome        <folder>            Java bin dir (overwrites PATH)
  -javaOption      <java option>       JVM option parameter, e.g. '-agentlib:...'
  -preHook         <programpath>       A program which will be executed before a
                                       suite run (can be added multiple times)
  -postHook        <programpath>       A program which will be executed after a
                                       suite run (can be added multiple times)
  -D               <JVM option>        JVM option to set a property at runtime,
                                       overwrites file based properties
  -browser         <browser>           Browser for the test execution
                                       (default: Firefox)
  -sahiHome        <folder>            Sahi installation folder

  -masterkey       <base64 AES key>    AES base64 key used by command 'encrypt'
                                       (default: use env var 'SAKULI_ENCRYPTION_KEY'
  -interface       <interface>         Network interface card name, used by
                                       command 'encrypt' as salt
                                       (default: 'auto' for use defeault NIC)

  -examples                            CLI usage examples
  -version                             Version info
  -help                                This help text

`, time.Now().Year())
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
	myFlagSet.StringVar(&sahiHome, "sahi_home", "", "Sahi installation folder")
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
			input.PrintExampleUsage()
		}
		detError := ""
		if len(os.Args) == 2 {
			detError += "ARGUMENT is missing specify one: "
		}
		input.ExitWithHelp("\n" + detError + "Only 'sakuli COMMAND ARGUMENT [OPTIONS]' is allowed, given: " + fmt.Sprint(os.Args))
	}

	sakuliProperties := map[string]string{"sakuli_home": helper.GetSakuliHome()}
	typ, argument := input.ParseArgs(append(os.Args[1:3], myFlagSet.Args()...))
	switch typ {
	case input.RunMode:
		input.TestRun(argument)
		sakuliProperties[input.RunMode] = argument
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
		sakuliProperties["browser"] = browser
	}
	if inter != "" {
		sakuliProperties["interface"] = inter
	}
	if masterkey != "" {
		sakuliProperties["masterkey"] = masterkey
	}
	if sahiHome != "" {
		sakuliProperties["sahiHome"] = sahiHome
	}
	joinedSakuliProperties := genSakuliPropertiesList(sakuliProperties)

	if len(preHooks) > 0 {
		fmt.Println("=========================== Starting Pre-Hooks =================================")
		for _, pre := range preHooks {
			execute.RunHandler(pre)
		}
		fmt.Println("=========================== Finished Pre-Hooks =================================")
	}

	sakuliReturnCode := execute.RunSakuli(javaExecutable, sakuliJars, javaOptions, javaProperties, joinedSakuliProperties)
	for loop > 0 {
		fmt.Printf("*** Loop mode - sleeping for %d seconds... ***\n", loop)
		time.Sleep(time.Duration(loop) * time.Second)
		execute.RunSakuli(javaExecutable, sakuliJars, javaOptions, javaProperties, joinedSakuliProperties)
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

func genSakuliPropertiesList(properties map[string]string) input.StringSlice {
	propertiesString := []string{}
	for k, v := range properties {
		propertiesString = append(propertiesString, fmt.Sprintf("--%s", k))
		propertiesString = append(propertiesString, v)
	}
	return propertiesString
}
