package worker

import (
	"fmt"
	"go-worker-pool/src/task"
	"go-worker-pool/src/utils"

	log "github.com/sirupsen/logrus"
)

type Worker struct {
	Id       int
	handler  utils.Functor
	taskChan chan *task.Task
	quit     chan bool
}

func NewWorker(id int, handler utils.Functor, taskChan chan *task.Task, quit chan bool) *Worker {
	return &Worker{Id: id, handler: handler, taskChan: taskChan, quit: quit}
}

func (wr *Worker) RunBackground() {
	log.Debug(fmt.Sprintf("Running worker %v", wr.Id))

	for {
		select {
		case task := <-wr.taskChan:
			wr.Process(task)
		case <-wr.quit:
			return
		}
	}

}

func (wr *Worker) Process(task *task.Task) utils.FunctorOutputData {
	log.Debug(fmt.Sprintf("Worker %v processing task %v", wr.Id, task.Id))
	return wr.handler(task.Data)
}

func (wr *Worker) Stop() {
	log.Debug(fmt.Sprintf("Stop worker %v", wr.Id))
	go func() {
		wr.quit <- true
	}()
}
