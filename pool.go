package tasker

import (
	"time"
)

// Pool ...
type Pool struct {
	workers      []*worker
	tasks        chan Task
	nextWorkerID int
}

// CreatePool create new worker pool
func CreatePool() *Pool {
	return &Pool{
		tasks:   make(chan Task, 4),
		workers: make([]*worker, 0),
	}
}

// AddWorkers ...
func (p *Pool) AddWorkers(count int) {
	for i := 0; i < count; i++ {
		p.AddWorker()
	}
}

// AddWorker ...
func (p *Pool) AddWorker() {
	p.nextWorkerID++
	p.workers = append(p.workers, &worker{
		tasks: p.tasks,
		id:    p.nextWorkerID,
	})
}

// Run pool
func (p *Pool) Run() {
	for _, w := range p.workers {
		go w.run()
	}

	for {
		time.Sleep(time.Minute)
		p.restartDownWorkers()
	}
}

func (p *Pool) restartDownWorkers() {
	for _, w := range p.workers {
		if w.isDown() {
			go w.run()
		}
	}
}

// Dispatch ...
func (p *Pool) Dispatch(t Task) {
	go func() {
		p.tasks <- t
	}()
}
