package persons

import "fmt"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonProvider interface {
	GetPersons() []Person
}

func GetAndPrintPersonsNames(provider PersonProvider) {
	p := provider.GetPersons()
	for _, n := range p {
		fmt.Println(n.Name)
	}
}
