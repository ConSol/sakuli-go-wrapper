package execute

import (
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"github.com/ConSol/sakuli-go-wrapper/input"
	"path/filepath"
	"strings"
	"os"
)

//RunSakuli starts the sakuli jar with javaProperties and sakuliProperties
func RunSakuli(javaExecutable, sakuliJars string, javaProperties, sakuliProperties input.StringSlice) int {
	classpath := helper.GenClassPath(filepath.Join(sakuliJars, "sakuli.jar"), filepath.Join(sakuliJars, "*"))
	args := []string{}
	args = append(args, javaProperties...)
	args = append(args, "-Duser.dir=" + os.Getenv("SAKULI_HOME") )
	args = append(args, "-classpath")
	args = append(args, classpath)
	args = append(args, "org.sakuli.starter.SakuliStarter")
	args = append(args, sakuliProperties...)
	fmt.Print("=========== Calling Sakuli JAR: ")
	fmt.Print(javaExecutable, " ", strings.Join(args, " "))
	fmt.Println(" ===========")
	returnCode, err := Execute(javaExecutable, args...)
	if err != nil {
		fmt.Println("Error while calling Sakuli JAR")
	}
	fmt.Printf("=========== Sakuli JAR finished with returncode: %d ===========\n", returnCode)
	return returnCode
}
