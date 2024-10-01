package even_number

import "errors"

type EvenNumber struct {
	n int
}

func New(n int) (*EvenNumber, error) {
	if n%2 == 0 {
		return &EvenNumber{n: n}, nil
	} else {
		return nil, errors.New("it not even")
	}
}

func (n EvenNumber) Set(k int) error {
	if k%2 == 0 {
		n.n = k
		return nil
	} else {
		return errors.New("it not even")
	}
}

func SetFn(en EvenNumber, k int) {
	en.n = k
}
