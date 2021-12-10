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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// newFileCmd represents the newFile command
var newFileCmd = &cobra.Command{
	Use:   "newFile",
	Short: "Creates a new file",
	Run: func(cmd *cobra.Command, args []string) {
		newFile(args)
	},
}

func init() {
	rootCmd.AddCommand(newFileCmd)
}

func newFile(name []string){
	f, err := os.OpenFile(filepath.Join(strings.Join(name, " ")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}


