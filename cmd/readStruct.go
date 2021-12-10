package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// readStructCmd represents the readStruct command
var readStructCmd = &cobra.Command{
	Use:   "readStruct",
	Short: "Searches the file for the requested Data and print it.",
	Run: func(cmd *cobra.Command, args []string) {
		readStruct(args[0], args[1:])
	},
}

func init() {
	rootCmd.AddCommand(readStructCmd)
}

func readStruct(data1 string, name []string){
	var a []Data
	data, err := ioutil.ReadFile(filepath.Join(strings.Join(name, "")))
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(data), "\n")
	temp := list[1:len(list)-1]
	for _, j := range temp{
		str := strings.Split(j, "|")
		id, err := strconv.Atoi(str[0])
		if err != nil{
			log.Panic(err)
		}
		age, err := strconv.Atoi(str[1])
		if err != nil{
			log.Panic(err)
		}
		phone, err := strconv.Atoi(str[3])
		if err != nil{
			log.Panic(err)
		}
		a = append(a, Data{ID: id, Age: age, Name: str[2], Phone: phone})
	}
	for _, j := range a{
		if j.Name == data1{
			fmt.Println(j.ID, j.Age, j.Name, j.Phone)
			return
		}
	}
	log.Fatal("Could not find the requested data!")
}