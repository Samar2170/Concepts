package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var GlobalKVStore = make(map[string]string)

// // transactions are started, committed, and rolled back
// type Transaction struct {
// 	store map[string]string
// }

// // NewTransaction creates a new transaction with BEGIN command
// func NewTransaction() *Transaction {
// 	transaction := Transaction{store: make(map[string]string)}
// 	return &transaction
// }

// type TransactionStack struct {
// 	stack []*Transaction
// }

// // NewTransactionStack creates a new transaction stack

// func NewTransactionStack() *TransactionStack {
// 	transactionStack := TransactionStack{}
// 	return &transactionStack
// }

// func (t *TransactionStack) Peek() (*map[string]string, bool) {
// 	if len(t.stack) == 0 {
// 		return nil, false
// 	}
// 	return &t.stack[len(t.stack)-1].store, true
// }
// func (t *TransactionStack) Commit() {
// 	for _, transaction := range t.stack {
// 		for k, v := range transaction.store {
// 			GlobalKVStore[k] = v
// 		}
// 	}
// 	t = NewTransactionStack()

// }

// func (t *TransactionStack) Push(transaction Transaction) {

// 	t.stack = append(t.stack, &transaction)
// }

// func (t *TransactionStack) Pop() {
// 	if len(t.stack) == 0 {
// 		return
// 	}
// 	t.stack = t.stack[:len(t.stack)-1]
// }

// func (t *TransactionStack) Size() int {
// 	return len(t.stack)
// }

// func run() {
// 	input := bufio.NewReader(os.Stdin)
// 	items := &TransactionStack{}

// 	for {
// 		fmt.Printf("Enter command: ")
// 		command, _ := input.ReadString('\n')
// 		switch command {
// 		case "BEGIN":
// 			transaction := NewTransaction()
// 			items.Push(*transaction)
// 		case "COMMIT":
// 			if items.Size() == 0 {
// 				fmt.Println("NO TRANSACTION")
// 				continue
// 			}
// 			items.Commit()
// 		case "ROLLBACK":
// 			if items.Size() == 0 {
// 				fmt.Println("NO TRANSACTION")
// 				continue
// 			}
// 			items.Pop()
// 		case "SET":

// 		case "END":
// 			return
// 		default:
// 			return

// 		}
// 	}
// }
