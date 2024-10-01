package persons_file_provider

import (
	"codewars/persons/persons"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type PersonsFileReader struct {
	FileName string
}

func (p PersonsFileReader) GetPersons() []persons.Person {
	jsonFile, err := os.Open(p.FileName)
	if err != nil {
		log.Fatal(err)
	}

	person := make([]persons.Person, 10)

	jsonReaded, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonReaded, &person)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range person {
		fmt.Println(n.Name)
	}
	fmt.Println(len(person))
	return nil
}
