package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type developer struct {
	timestamp    string
	name         string
	email        string
	gitLink      string
	phoneNumber  string
	linkedInLink string
	websiteLink  string
	// Photo        string // unused yet, fetch from his github instead (if any)
	comments string
}

func main() {
	csvFile, _ := os.Open("Majalengka Developer Contact Information.csv")
	reader := csv.NewReader(csvFile)
	var developers []developer
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		developers = append(developers, developer{
			timestamp:    line[0],
			name:         line[1],
			email:        line[2],
			gitLink:      line[3],
			phoneNumber:  line[4],
			linkedInLink: line[5],
			websiteLink:  line[6],
			comments:     line[8],
		})
	}
	fmt.Println(developers)
	// TODO Print out HTML file based on this data
	a := gomod.Document.GetElementById("something")
	fmt.Printf("\n%#v\n", a)
}
