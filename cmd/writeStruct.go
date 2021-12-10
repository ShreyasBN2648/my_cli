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
	"strconv"

	"github.com/spf13/cobra"
)

type Data struct {
	ID    int
	Age   int
	Name  string
	Phone int
}

// writeStructCmd represents the writeStruct command
var writeStructCmd = &cobra.Command{
	Use:   "writeStruct",
	Short: "Writes struct data into the file",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[1])
		if err != nil{
			log.Panic(err)
		}
		age, err := strconv.Atoi(args[2])
		if err != nil{
			log.Panic(err)
		}
		phone, err := strconv.Atoi(args[4])
		if err != nil{
			log.Panic(err)
		}
		writeStruct(args[0], id, age, args[3], phone)
	},
}

func init() {
	rootCmd.AddCommand(writeStructCmd)
}

func writeStruct(name string, ID, Age int, Name string, Phone int) {
	var a Data = Data{
		ID:    ID,
		Age:   Age,
		Name:  Name,
		Phone: Phone,
	}
	f, err := os.OpenFile(filepath.Join(name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%v|%v|%v|%v\n", a.ID, a.Age, a.Name, a.Phone)
}
