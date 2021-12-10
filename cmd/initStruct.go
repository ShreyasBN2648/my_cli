/*
Package cmd
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// initStructCmd represents the initStruct command
var initStructCmd = &cobra.Command{
	Use:   "initStruct",
	Short: "Creates and initializes the file with the Struct format",
	Run: func(cmd *cobra.Command, args []string) {
		initStruct(args)
	},
}

func init() {
	rootCmd.AddCommand(initStructCmd)
}

func initStruct(name []string){
	f, err := os.OpenFile(filepath.Join(strings.Join(name, "")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%v|%v|%v|%v\n", "ID", "Age", "Name", "Phone")
}