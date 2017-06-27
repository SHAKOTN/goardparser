package utils

import (
	"log"
	"sync"
)

type Task struct {
	Err error

	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}

func (t *Task) Run(wg *sync.WaitGroup) {
	t.Err = t.f()
	wg.Done()
}

type Pool struct {
	Tasks []*Task

	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

func (p *Pool) HasErrors() bool {
	for _, task := range p.Tasks {
		if task.Err != nil {
			return true
		}
	}
	return false
}

func (p *Pool) Run() {
	log.Printf("Running %v task(s) at concurrency %v.",
		len(p.Tasks), p.concurrency)

	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	// close all workers
	close(p.tasksChan)

	p.wg.Wait()
}

func (p *Pool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}