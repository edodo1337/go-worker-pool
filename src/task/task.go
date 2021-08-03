package task

import (
	"go-worker-pool/src/utils"
)

type Task struct {
	Id   int
	Data utils.FunctorInputData
}

func NewTask(id int, data utils.FunctorInputData) *Task {
	return &Task{
		Id:   id,
		Data: data,
	}
}
