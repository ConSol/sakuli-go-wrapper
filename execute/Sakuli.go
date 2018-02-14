package execute

import (
	"fmt"
	"github.com/ConSol/sakuli-go-wrapper/helper"
	"github.com/ConSol/sakuli-go-wrapper/input"
	"path/filepath"
	"strings"
)

//RunSakuli starts the sakuli jar with javaProperties and sakuliProperties
func RunSakuli(javaExecutable string, sakuliJars string, javaOptions input.StringSlice, javaProperties input.StringSlice, sakuliProperties map[string]string) int {
	classpath := helper.GenClassPath(filepath.Join(sakuliJars, "sakuli.jar"), filepath.Join(sakuliJars, "*"))
	args := []string{}
	args = append(args, javaOptions...)
	args = append(args, javaProperties...)
	args = append(args, "-classpath")
	args = append(args, classpath)
	args = append(args, "org.sakuli.starter.SakuliStarter")
	args = append(args, genSakuliRunPropertiesList(sakuliProperties)...)
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

func genSakuliRunPropertiesList(properties map[string]string) input.StringSlice {
	propertiesString := []string{}
	for k, v := range properties {
		propertiesString = append(propertiesString, fmt.Sprintf("--%s", k))
		propertiesString = append(propertiesString, v)
	}
	return propertiesString
}

//RunSakuliUI starts the sakuli UI jar with parsed javaProperties
func RunSakuliUI(javaExecutable string, sakuliJars string, javaOptions input.StringSlice, javaProperties input.StringSlice, sakuliProperties map[string]string) int {
	jarName := helper.GenClassPath(filepath.Join(sakuliJars, "sakuli-ui-web.jar"))
	javaProperties = appendUiProps(javaProperties, sakuliProperties)
	args := []string{}
	args = append(args, javaOptions...)
	args = append(args, javaProperties...)
	args = append(args, "-jar")
	args = append(args, jarName)
	fmt.Println("============================== Calling Sakuli UI JAR ===========================")
	fmt.Println("command:", javaExecutable, strings.Join(args, " "))
	fmt.Println("")
	returnCode, err := Execute(javaExecutable, args...)
	if returnCode != 0 {
		if err != nil {
			fmt.Println("Error while calling Sakuli UI JAR:\n" + err.Error())
		}
		fmt.Printf("=================== Sakuli UI JAR finished with returncode: %d =================\n", returnCode)
	}
	return returnCode
}

func appendUiProps(javaProps input.StringSlice, sakuliProperties map[string]string) input.StringSlice {
	const SAKULI_HOME_FOLDER = "sakuli.home.folder"
	const SAKULI_UI_ROOT_DIRECTORY = "sakuli.ui.root.directory"
	javaProps = append(javaProps, "-D"+SAKULI_HOME_FOLDER+"="+sakuliProperties[input.OptSakuliHome])
	javaProps = append(javaProps, "-D"+SAKULI_UI_ROOT_DIRECTORY+"="+sakuliProperties[input.UiMode])
	return javaProps
}
