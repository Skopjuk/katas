package i_hope_the_last_one

import (
	"fmt"
	"log"
	"os"
)

type Writer interface {
	Write(str string)
}

type Console struct {
	str string
}

func NewConsole(s string) Console {
	return Console{str: s}
}

type FileLog struct {
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printValues(w Writer, s []string) {
	for i := range s {
		w.Write(s[i])
	}

}

func (c *Console) Write(str string) {
	fmt.Println(str)
}

func (f *FileLog) Write(str string) {
	file, err := os.OpenFile("i_hope_the_last_one/log_file.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	str = str + "\n"
	_, err = file.WriteString(str)
	check(err)
}

func PrintValuesMain() {
	w := FileLog{}
	var s []string
	s = append(s, " good morning everybody Boris")
	s = append(s, " goodbye Boris")
	printValues(&w, s)
}
