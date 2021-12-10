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

// newFolderCmd represents the newFolder command
var newFolderCmd = &cobra.Command{
	Use:   "newFolder",
	Short: "Creates a new folder",
	Run: func(cmd *cobra.Command, args []string) {
		newFolder(args)
	},
}

func init() {
	rootCmd.AddCommand(newFolderCmd)
}

func newFolder(name []string){
	err := os.Mkdir(filepath.Join(strings.Join(name," ")), 0755)
	if err != nil {
		log.Fatal(err)
	}
}