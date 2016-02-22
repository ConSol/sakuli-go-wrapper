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
	var preHooks input.StringSlice
	var postHooks input.StringSlice
	var browser string
	var encrypt string
	var inter string
	var run string
	var sahiHome string
	var version bool

	sakuliJars := filepath.Join(os.Getenv("SAKULI_HOME"), "libs", "java")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Generic Sakuli test starter.
%d - The Sakuli team / Philip Griesbacher.
http://www.sakuli.org
https://github.com/ConSol/sakuli

Usage:   sakuli[.exe] COMMAND [OPTIONS]
         sakuli -help
         sakuli -version
         sakuli -run <sakuli suite> [OPTIONS]
         sakuli -encrypt <secret> [OPTIONS]

Commands:
         run
         encrypt

Options:
         -loop <minutes>           Loop this suite, wait n seconds between
                                   executions, 0 means no loops (default: 0)
         -javaHome <folder>        Java bin dir (overrides PATH)
         -preHooks <programpath>   A programm which will be executed before
                                   sakuli (Can be added multiple times)
         -postHooks  <programpath> A programm which will be executed after
                                   sakuli (Can be added multiple times)
         -D <JVM option>           JVM option to set a property on runtime,
                                   overrides the 'sakuli.properties'
         -browser <browser>        Browser for the test execution (default: Firefox)
         -interface <interface>    Network interface used for encryption
         -sahiHome <folder>        Sahi installation folder
         -version                  Version info


`, time.Now().Year())
	}

	flag.IntVar(&loop, "loop", 0, "loop this suite, wait n seconds between executions, 0 means no loops (default: 0)")
	flag.StringVar(&javaHome, "javahome", "", "Java bin dir (overrides PATH)")
	flag.Var(&preHooks, "preHook", "A programm which will be executed before sakuli (Can be added multiple times)")
	flag.Var(&postHooks, "postHook", "A programm which will be executed after sakuli (Can be added multiple times)")

	flag.StringVar(&encrypt, "encrypt", "", "encrypt a secret")
	flag.StringVar(&run, "run", "", "run a sakuli test suite")

	flag.Var(&javaProperties, "D", "JVM option to set a property on runtime, overrides the 'sakuli.properties'")
	flag.StringVar(&browser, "browser", "", "browser for the test execution (default: Firefox)")
	flag.StringVar(&inter, "interface", "", "network interface used for encryption")
	flag.StringVar(&sahiHome, "sahi_home", "", "Sahi installation folder")
	flag.BoolVar(&version, "version", false, "version info")
	flag.Parse()

	if version {
		input.PrintVersion()
	}
	//input.TestRun(run)

	javaExecutable := input.TestJavaHome(javaHome)
	javaProperties = javaProperties.AddPrefix("-D")
	sakuliProperties := map[string]string{"sakuli_home": helper.GetSahiHome()}

	if browser != "" {
		sakuliProperties["browser"] = browser
	}
	if inter != "" {
		sakuliProperties["interface"] = inter
	}
	if encrypt != "" {
		sakuliProperties["encrypt"] = encrypt
	}
	if run != "" {
		sakuliProperties["run"] = run
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

	sakuliReturnCode := execute.RunSakuli(javaExecutable, sakuliJars, javaProperties, joinedSakuliProperties)
	for loop > 0 {
		fmt.Printf("*** Loop mode - sleeping for %d seconds... ***\n", loop)
		time.Sleep(time.Duration(loop) * time.Second)
		execute.RunSakuli(javaExecutable, sakuliJars, javaProperties, joinedSakuliProperties)
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
