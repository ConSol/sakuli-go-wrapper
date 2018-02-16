package input

import (
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func PrintHelp(uiInstalled bool) {
	uiTextSmall, uiTextLong := genUiHelpText(uiInstalled)

	fmt.Fprintf(os.Stderr, `Generic Sakuli test starter.
%d - The Sakuli team <sakuli@consol.de>
http://www.sakuli.org
https://github.com/ConSol/sakuli

Usage: sakuli[.exe] COMMAND ARGUMENT [OPTIONS]

  sakuli -help
  sakuli -examples
  sakuli -version
  sakuli run       <sakuli suite path> [OPTIONS]%s
  sakuli encrypt   <secret> [OPTIONS]
  sakuli create    <object> [OPTIONS]

Commands:
  run              <sakuli suite path>  runs a Sakuli test suite%s                
  encrypt          <secret>             encrypting a secret                    
  create           <object>             create different objects               
                                                                               
Objects:                                                                       
  masterkey                             Base64 decoded AES 128-bit key
                                        (for environment based encryption)
                                                                               
Options:                                                                       
  -loop            <seconds>            Loop this suite, wait n seconds between
                                        executions, 0 means no loops (default)
  -javaHome        <folder>             Java bin dir (overwrites PATH)         
  -javaOption      <java option>        JVM option parameter, e.g. '-agentlib:'
  -preHook         <programpath>        Program which will be executed before a
                                        suite run (can be added multiple times)
  -postHook        <programpath>        Program which will be executed after a
                                        suite run (can be added multiple times)
  -D               <JVM option>         JVM option to set a property at runtime
                                        (overwrites file based properties)       
  -browser         <browser>            Browser for the test execution         
                                        (default: Firefox)                     
  -sahiHome        <folder>             Sahi installation folder 

  -masterkey       <base64 AES key>     AES base64 key used by 'sakuli encrypt'
                                        (default: use environment variable 
                                        'SAKULI_ENCRYPTION_KEY')
  -interface       <interface>          Network interface card name, used by   
                                        'sakuli encrypt' as salt              
                                        (default: 'auto' for use defeault NIC) 
                                                                               
  -examples                             CLI usage examples                     
  -version                              Version info                           
  -help                                 This help text                         

`, time.Now().Year(), uiTextSmall, uiTextLong)
}

//if ui is not installed don't show the UI option in the starter
func genUiHelpText(uiInstalled bool) (string, string) {
	if uiInstalled {
		return `
  sakuli ui        <root context path> [OPTIONS]`, `
  ui               <root context path>  starts the Sakuli UI with access to the
                                        given root path`
	}
	return "", ""
}

//PrintVersion prints the sakuli version and env variables.
func PrintVersion() {
	versionFile := filepath.Join(helper.GetSakuliHome(), "bin", "resources", "version.txt")
	data, err := ioutil.ReadFile(versionFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Print("\n--- Environment variables ---")
	fmt.Printf(`
SAKULI_ROOT:                  %s
SAKULI_HOME:                  %s
MOZ_DISABLE_OOP_PLUGINS:      %s
MOZ_DISABLE_AUTO_SAFE_MODE:   %s
MOZ_DISABLE_SAFE_MODE_KEY:    %s
`, helper.GetSakuliRoot(), helper.GetSakuliHome(), os.Getenv("MOZ_DISABLE_OOP_PLUGINS"), os.Getenv("MOZ_DISABLE_AUTO_SAFE_MODE"), os.Getenv("MOZ_DISABLE_SAFE_MODE_KEY"))
	os.Exit(0)
}

//if ui is not installed don't show the UI examples
func PrintExampleUsage(uiInstalled bool) {
	fmt.Printf(`Sakuli CLI Examples:

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
%s
Encrypt secrets:
▶ Default mode: Encrypt a secret using  master key from
  environment var 'SAKULI_ENCRYPTION_KEY':
    export SAKULI_ENCRYPTION_KEY=Bsqs/IR1jW+eibNrdYvlAQ==
    sakuli encrypt topsecret
▶ Encrypt a secret using an provided masterkey:
    sakuli encrypt topsecret -masterkey Bsqs/IR1jW+eibNrdYvlAQ==
▶ Create a new random master key masterkey Base64 AES-128 key:
    sakuli create masterkey
▶ Encrypt a secret using an automatic determend NIC as salt:
    sakuli encrypt topsecret -interface auto
▶ Encrypt a secret using eth0 as salt NIC:
    sakuli encrypt topsecret -interface eth0
▶ Show interfaces available for encryption:
    sakuli encrypt topsecret -interface list

Others:
▶ Show version (use this information when submitting bugs):
    sakuli -version
`, genUiExampleText(uiInstalled))
	os.Exit(0)
}

func genUiExampleText(uiInstalled bool) string {
	if uiInstalled {
		return `
Start Sakuli UI:
(default user: admin, password: sakuli123)
▶ Start UI at given path as root context path
    sakuli ui ~/sakuli/example_test_suites
▶ Change default user and password
    sakuli ui . -D security.default-username=myuser -D security.default-password=mypassword
▶ Disable authentication
    sakuli ui . -D app.authentication.enabled=true
`
	}
	return ``
}
