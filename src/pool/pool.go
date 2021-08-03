package pool

import (
	"go-worker-pool/src/task"
	"go-worker-pool/src/utils"
	"go-worker-pool/src/worker"
)

type Pool struct {
	workers     []*worker.Worker
	concurrency int
	handler     utils.Functor
	collector   chan *task.Task

	quit chan bool
}

func NewPool(concurrency int, handler utils.Functor) *Pool {
	quit := make(chan bool)
	collector := make(chan *task.Task)

	var workers []*worker.Worker
	for i := 0; i < concurrency; i++ {
		workers = append(workers, worker.NewWorker(i, handler, collector, quit))
	}

	return &Pool{concurrency: concurrency, workers: workers, handler: handler, quit: quit, collector: collector}
}

func (p *Pool) Process(tasks []*task.Task) {
	for _, wr := range p.workers {
		go wr.RunBackground()
	}
	for _, task := range tasks {
		p.collector <- task
	}
}
