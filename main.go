package main

import (
	"encoding/json"
	"fmt"
	"github.com/grahambrooks/imprint/info"
	"github.com/spf13/cobra"
	"os"
	"reflect"
)

var infoFlagValue bool
var xDefsFlagValue bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&infoFlagValue, "info", "i", false, "info")
	rootCmd.PersistentFlags().BoolVarP(&xDefsFlagValue, "x-defs", "x", false, "lists x-defs mapping for Bazel builds")
}

var rootCmd = &cobra.Command{
	Use:   "imprint",
	Short: "Version and meta-data imprinting tool for Go and Bazel",
	Example: `
To generate a space separated key value list of relevant environment variable entries

	imprint

Use this command to imprint the values into your binary that links with imprint:

	bazel run --stamp --workspace_status_command=imprint --verbose_failures //:target 

imprint --x-defs

Outputs the x-defs value for bazel build files. This mapping is built into the imprint library.

imprint --info

Outputs a JSON formatted rendering of the default build structures. If all goes well the values printed reflect the environment variable values at the time the application was built.

`,
	Run: func(cmd *cobra.Command, args []string) {
		if infoFlagValue {
			encoder := json.NewEncoder(os.Stdout)
			_ = encoder.Encode(info.Block())
		} else {
			if xDefsFlagValue {
				fmt.Printf("    x_defs = {\n")
				for key, value := range info.DefinitionMap {
					fmt.Printf("      \"%s.%s\": \"{%s}\",\n", info.PackagePath, value.Name, key)
				}
				fmt.Printf("    },\n")
			} else {
				info := info.InfoBlock{}

				printImprintingEnvironmentVariables(&info)

				for _, name := range EnvironmentVariables {
					if len(os.Getenv(name)) > 0 {
						fmt.Printf("%s %s\n", name, os.Getenv(name))
					}
				}
			}
		}
		// Do Stuff Here
	},
}

func printImprintingEnvironmentVariables(s interface{}) {
	reflectType := reflect.TypeOf(s).Elem()
	reflectValue := reflect.ValueOf(s).Elem()

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		//typeName := field.Name
		//
		//valueType := reflectValue.Field(i).Type()

		switch reflectValue.Field(i).Kind() {
		case reflect.String:
			imprintTag := field.Tag.Get("imprint")
			fmt.Printf("%s %s \n", field.Name, imprintTag)
		case reflect.Struct:
			valueValue := reflectValue.Field(i)
			printImprintingEnvironmentVariables(&valueValue)
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
