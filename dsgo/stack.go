package main

type Stack []int

func (s *Stack) Len() int { return len(*s) }

func (s *Stack) Append(x int) {
	*s = append(*s, x)
}
func (s *Stack) Pop() int {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

type Transaction map[string]string

type TransactionStack []Transaction

func (s *TransactionStack) Len() int { return len(*s) }

func (s *TransactionStack) Append(x Transaction) {
	*s = append(*s, x)
}
func (s *TransactionStack) Pop() Transaction {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}
