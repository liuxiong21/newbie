package chart01

import (
	"errors"
)

/*structs Stacker{
	top int
	elements []interface{}
	Push(em interface{})
	Pop() interface{}
}*/

type Stacker []interface{}

func (stack Stacker) Len() int {
	return len(stack)
}

func (stack Stacker) Cap() int {
	return cap(stack)
}

func (stack *Stacker) Push(em interface{}) {
	*stack = append(*stack, em)
}

/**
func (stack Stacker) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Stacker is empty")
	}
	return stack[len(stack)-1], nil
}
*/

func (stack *Stacker) Top() (interface{}, error) {
	if len(*stack) == 0 {
		return nil, errors.New("Stacker is empty")
	}
	var val = *stack
	return val[len(*stack)-1], nil
}
