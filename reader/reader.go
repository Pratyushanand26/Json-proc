package reader

import (
	"encoding/json"
	"log"
	"os"
)

type RawJson struct {
	id     int    `json:"id"`
	name   string `json:"name"`
	score  int    `json:"score"`
	status string `json:"status"`
}

func Read(out chan<- RawJson, inpath string) {
	filecontent, error := os.ReadFile(inpath)

	if error != nil {
		log.Fatal(error)
	}

	var records []RawJson

	if err := json.Unmarshal(filecontent, &records); err != nil {
		log.Fatal("failed to unmarshal", err)
	}

	for _, record := range records {
		out <- record
	}

}
