package main

import "log"

type Job interface {
	Do()
}

type Worker struct {
	TaskQueue chan Job
}

type AddJob struct {
	Num1 int
	Num2 int
}

func (j AddJob) Do() {
	log.Println(j.Num1 + j.Num2)
}

type SubtractJob struct {
	Num1 int
	Num2 int
}

func testIt() {
	// worker := Worker{TaskQueue: taskQueue}

}
