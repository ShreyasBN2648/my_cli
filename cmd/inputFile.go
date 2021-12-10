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
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

// inputFileCmd represents the inputFile command
var inputFileCmd = &cobra.Command{
	Use:   "inputFile",
	Short: "Adds data into the file",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(inputFileCmd)
}

func inputFile(name string, data string){
	f, err := os.OpenFile(filepath.Join(name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%v\n", data)
}
