package inmemory_provider

import "codewars/persons/persons"

type InMemoryProvider struct {
	P []persons.Person
}

func (i InMemoryProvider) GetPersons() []persons.Person {
	return i.P
}
