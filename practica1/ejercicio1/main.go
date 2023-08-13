package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var filename string

func init(){
  flag.StringVar(&filename,"filename","","Text file to read")
  flag.Parse()
}

func textParser(){
  file, err :=
  os.Open(filename)
  if err != nil{
    log.Fatal(err)
  }
  defer file.Close()
  
  text, err := io.ReadAll(file)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Number of characters: ",len(text))
  fmt.Println("Number: of lines",strings.Count(string(text), "\n"))
  fmt.Println("Number of words: ",len(strings.Fields(string(text)))+1)

}

func main(){
  if filename == ""{
    panic(errors.New("Try with :!go run main.go -filename <nombre archivo>"))
  }
  textParser()
 }
