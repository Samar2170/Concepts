package main

import (
	"log"
	"runtime"
	"time"
)

//patterns
// 1. 1 goroutine
// 2. 2 goroutines
// 3. 2 goroutines with a channel
// 4. 2 goroutines with a channel and a waitgroup
// 5. 2 goroutines with a channel and a waitgroup and a mutex

// Fan in Fan out
// Generator
// Pipeline
// Job Queue
// Semaphore
// Worker Pool

type T = interface{}

type WorkerPool interface {
	Run()
	AddTask(task func())
}
type workerPool struct {
	maxWorkers  int
	queuedTaskC chan func()
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorkers; i++ {
		go func(workerId int) {
			for task := range wp.queuedTaskC {
				task()
			}
		}(i + 1)
	}
}
func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}

func testConcurrency() {
	waitC := make(chan bool)
	go func() {
		for {
			log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	totalWorkers := 3
	wp := &workerPool{maxWorkers: totalWorkers, queuedTaskC: make(chan func())}
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTask := 5
	resultC := make(chan result, totalTask)
	for i := 0; i < totalTask; i++ {
		taskId := i
		wp.AddTask(func() {
			time.Sleep(2 * time.Second)
			resultC <- result{id: taskId, value: taskId * taskId}
		})
	}

	for i := 0; i < totalTask; i++ {
		r := <-resultC
		log.Printf("[main] result: %v", r)
	}

	<-waitC
}
