package main

import (
	"fmt"
	"os"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

var (
	apiClient = client.NewAPIClient(
		settings.NewAuthenticationSettings(
			os.Getenv("KEY"),
			os.Getenv("SECRET"),
		),
		settings.NewHttpSettings(
			os.Getenv("CONDUCTOR_SERVER_URL"),
		))
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
	metadataClient   = client.MetadataResourceApiService{APIClient: apiClient}
)

func main() {

}

func StartWorkers() {
	taskRunner.StartWorker("task1", Task1, 1, time.Millisecond*100)
	taskRunner.StartWorker("task2", Task2, 1, time.Millisecond*100)
}

func Task1(task *model.Task) (interface{}, error) {

	//To fail the task send an error
	//return nil, errors.New("bad idea")

	return map[string]interface{}{
		"greetings": "Hello, " + fmt.Sprintf("%v", task.InputData["name"]),
	}, nil
}

func Task2(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"zip": "10121",
	}, nil
}
