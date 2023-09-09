package main

import (
	"fmt"
)

var GlobalStore = make(map[string]string)

type Transaction struct {
	store map[string]string
	next  *Transaction
}

type TransactionStack struct {
	top  *Transaction
	size int
}

func (ts *TransactionStack) Peek() *Transaction {
	return ts.top
}

func (ts *TransactionStack) Commit() {
	ActiveTransaction := ts.Peek()
	if ActiveTransaction != nil {
		for k, v := range ActiveTransaction.store {
			GlobalStore[k] = v
			if ActiveTransaction.next != nil {
				ActiveTransaction.next.store[k] = v
			}
		}
	} else {
		fmt.Printf("No active transaction")
	}
}

func (ts *TransactionStack) Rollback() {
	if ts.top == nil {
		fmt.Printf("No active transaction")
	} else {
		for key := range ts.top.store {
			delete(ts.top.store, key)
		}
	}
}
func Get(key string, T *TransactionStack) {
	ActiveTransaction := T.Peek()
	if ActiveTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
			fmt.Println(val)
		} else {
			fmt.Println("NULL")
		}
	} else {
		if val, ok := ActiveTransaction.store[key]; ok {
			fmt.Println(val)
		} else {
			fmt.Println("NULL")
		}
	}
}

func Set(key string, value string, T *TransactionStack) {
	ActiveTransaction := T.Peek()
	if ActiveTransaction == nil {
		GlobalStore[key] = value
	} else {
		ActiveTransaction.store[key] = value
	}
}
