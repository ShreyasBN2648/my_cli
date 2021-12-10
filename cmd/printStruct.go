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

// printStructCmd represents the printStruct command
var printStructCmd = &cobra.Command{
	Use:   "printStruct",
	Short: "Print the data in the file as a struct",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("data")
		if format == "list" {
			printData(args)
		}else{
			printStruct(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(printStructCmd)
	printStructCmd.PersistentFlags().String("data", "", "A search term for a result format. 'list' or '' ")
}

func printStruct1(name []string, a *[]Data){
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
		*a = append(*a, Data{ID: id, Age: age, Name: str[2], Phone: phone})
	}
}

func printStruct(name []string)  {
	var a []Data
	printStruct1(name, &a)
	fmt.Println(a)
}

func printData(name []string){
	var a []Data
	printStruct1(name, &a)
		for _,j := range a{
			fmt.Println(j.ID, j.Age, j.Name, j.Phone)
		}

}
