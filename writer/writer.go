package writer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type RawJson struct {
	id     int
	name   string
	score  int
	status string
}

func Write(in <-chan RawJson) {

	dir := `D:\Projects\random\json-proc`

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("error during making directory", err)
	}

	rawdata:=<-in

	data,err:=json.Marshal(rawdata)
	if err!=nil{
		log.Fatal(err)
	}

	os.WriteFile(dir,data,0655);
}
