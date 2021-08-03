package main

import (
	"fmt"
	"go-worker-pool/src/pool"
	"go-worker-pool/src/task"
	"go-worker-pool/src/utils"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	handler := func(input utils.FunctorInputData) utils.FunctorOutputData {
		logrus.Debug(fmt.Sprintf("Message from input %v\n", input["message"]))
		output := make(map[string]interface{})
		time.Sleep(time.Millisecond * 100)
		output["message"] = "hello_response"
		return output
	}

	tasks := []*task.Task{
		{
			Id:   1,
			Data: map[string]interface{}{"message": "hello1"},
		},
		{
			Id:   2,
			Data: map[string]interface{}{"message": "hello2"},
		},
		{
			Id:   3,
			Data: map[string]interface{}{"message": "hello3"},
		},
		{
			Id:   4,
			Data: map[string]interface{}{"message": "hello4"},
		},
	}
	logrus.SetLevel(logrus.DebugLevel)
	pool := pool.NewPool(3, handler)
	pool.Process(tasks)

}
