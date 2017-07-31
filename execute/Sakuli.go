package execute

import (
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"github.com/ConSol/sakuli-go-wrapper/input"
	"path/filepath"
	"strings"
)

//RunSakuli starts the sakuli jar with javaProperties and sakuliProperties
func RunSakuli(javaExecutable, sakuliJars string, javaOptions, javaProperties, sakuliProperties input.StringSlice) int {
	classpath := helper.GenClassPath(filepath.Join(sakuliJars, "sakuli.jar"), filepath.Join(sakuliJars, "*"))
	args := []string{}
	args = append(args, javaOptions...)
	args = append(args, javaProperties...)
	args = append(args, "-classpath")
	args = append(args, classpath)
	args = append(args, "org.sakuli.starter.SakuliStarter")
	args = append(args, sakuliProperties...)
	fmt.Println("============================== Calling Sakuli JAR ==============================")
	fmt.Println("command:", javaExecutable, strings.Join(args, " "))
	fmt.Println("")
	returnCode, err := Execute(javaExecutable, args...)
	if returnCode < 0 || returnCode > 6 {
		if err != nil {
			fmt.Println("Error while calling Sakuli JAR:\n" + err.Error())
		}
		fmt.Printf("=================== Sakuli JAR finished with returncode: %d ====================\n", returnCode)
	}
	return returnCode
}
