package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
)


func main(){
var(
	Inpath=flag.String("in","","input json path")
	Outpath=flag.String("out","","output json path")
	Workers=flag.Int("Workers",1,"no of Workers")
)

flag.Parse()

if *Inpath=="" || *Outpath=="" {
 fmt.Println("inputpath or Outpath is empty")
 os.Exit(1)
}

if *Workers<0{
	fmt.Println("Workers should be greater than 0")
	os.Exit(1)
}

if _,err:=os.Stat(*Inpath); os.IsNotExist(err){
	fmt.Println("file doesnt exist")
	os.Exit(1)
}

if _,err:=os.Stat(*Outpath); os.IsNotExist(err){
	fmt.Println("file doesnt exist")
	os.Exit(1)
}

if !strings.HasSuffix(*Inpath,".json") && !strings.HasSuffix(*Outpath,".json") {
	fmt.Println("file is not of jsonformat")
	os.Exit(1)
}
}