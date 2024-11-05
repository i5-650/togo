package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []int

// Pop removes and returns the first element of the stack
func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	// Retrieve the first element
	element := (*s)[0]

	// Remove the first element by re-slicing
	*s = (*s)[1:]
	return element, nil
}

func ParseIds(ids string) (Stack, error) {
	elements := strings.Split(ids, ".")
	stack := make(Stack, len(elements))

	for i, elem := range elements {
		num, err := strconv.Atoi(elem)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int: %v", elem, err)
		}
		stack[i] = num - 1
	}
	return stack, nil
}
