package main

var x, y = 50, 1000

type AddJob struct {
	x, y int
}

func (j AddJob) Do() int {
	return j.x + j.y
}

type SubJob struct {
	x, y int
}

func (j SubJob) Do() int {
	return j.x - j.y
}

type MultiplicationJob struct {
	x, y int
}

func (j MultiplicationJob) Do() int {
	return j.x * j.y
}

type Job interface {
	Do() int
}

func main() {

}
