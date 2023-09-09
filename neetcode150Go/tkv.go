package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Transactionn struct {
	store map[string]string
	next  *Transactionn
}

type TransactionStackk struct {
	top  *Transactionn
	size uint
}

var GlobalKVStoree = make(map[string]string)

func NewTransactionStackk() *TransactionStackk {
	return &TransactionStackk{top: nil, size: 0}
}

func (t *TransactionStackk) Peek() *Transactionn {
	return t.top
}

func (t *TransactionStackk) Push() {
	tr := Transactionn{store: make(map[string]string)}
	if t.size == 0 {
		t.top = &tr
	}
	next := t.top.next
	t.top = &tr
	t.top.next = next
	t.size++
}
func (t *TransactionStackk) Pop() {
	if t.top == nil {
		return
	}
	t.top = t.top.next
	t.size--

}

func (t *TransactionStackk) RollBack() {
	if t.size == 0 {
		fmt.Println("No Transactions to rollback")
	}
	t.top = t.top.next
	fmt.Println("Transaction Rolled back")
}

func (t *TransactionStackk) Commit() {
	active := t.Peek()
	if active == nil {
		fmt.Println("Nothing in Active tr")
	} else {
		for k, v := range active.store {
			GlobalKVStoree[k] = v
		}

	}
}

func Set(key string, value string, t *TransactionStackk) {
	active := t.Peek()
	if active == nil {
		t.Push()
		active = t.Peek()
		active.store[key] = value
	} else {
		active.store[key] = value
	}

}

func Get(key string, t *TransactionStackk) {
	active := t.Peek()
	if active == nil {
		if val, ok := GlobalKVStoree[key]; ok {
			fmt.Println(val)
		} else {
			fmt.Println("Key not Set")
		}
	} else {
		if val, ok := active.store[key]; ok {
			fmt.Println(val)
		} else {
			fmt.Println("Key not Set")
		}
	}
}

func transact() {
	input := bufio.NewReader(os.Stdin)
	items := &TransactionStackk{}

	for {
		fmt.Printf("Enter command: ")
		command, _ := input.ReadString('\n')
		ops := strings.Fields(command)
		switch ops[0] {
		case "BEGIN":
			items.Push()
		case "COMMIT":
			if items.size == 0 {
				fmt.Println("NO TRANSACTION")
				continue
			}
			items.Commit()
		case "ROLLBACK":
			if items.size == 0 {
				fmt.Println("NO TRANSACTION")
				continue
			}
			items.Pop()
		case "SET":
			Set(ops[1], ops[2], items)
		case "GET":
			Get(ops[1], items)
		case "END":
			os.Exit(0)
		default:
			fmt.Println("ERROR: Unrecognised Operation", ops[0])

		}
	}
}
func testTKV() {
	items := NewTransactionStackk()
	Set("a", "1", items)
	Set("b", "2", items)
	Set("c", "3", items)
	items.Commit()
	Get("a", items)
	Get("d", items)
	fmt.Println("GlobalKVStoree", GlobalKVStoree)
}
