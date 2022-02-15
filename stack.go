package main

type Stack []interface{}

func NewStack(vals ...interface{}) *Stack {
	rv := Stack{}
	for _, v := range vals {
		rv = append(rv, v)
	}
	return &rv
}

func (s *Stack) Peek() interface{} {
	return (*s)[s.Size()-1]
}

func (s *Stack) Push(v interface{}) *Stack {
	(*s) = append(*s, v)
	return s
}

func (s *Stack) Pop() interface{} {
	rv := (*s)[s.Size()-1]
	(*s) = (*s)[:s.Size()-1]
	return rv
}

func (s *Stack) Size() int {
	return len(*s)
}
