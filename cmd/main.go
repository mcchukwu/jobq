package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID      string
	Payload string
}

var queue []Job

func main() {
	go Worker()

	for i := range 10 {
		Enqueue(Job{
			ID:      fmt.Sprintf("%d", i),
			Payload: "task",
		})
	}

	select {}
}

// Task queue / Message broker
func Enqueue(job Job) {
	queue = append(queue, job)
}

func Dequeue() (Job, bool) {
	if len(queue) == 0 {
		return Job{}, false
	}

	job := queue[0]
	queue = queue[1:]
	return job, true
}

// Worker
func Worker() {
	for {
		job, ok := Dequeue()
		if !ok {
			time.Sleep(1 * time.Second)
			continue
		}

		Process(job)
	}
}

func Process(job Job) {
	fmt.Println("Processing:", job.Payload)
}
