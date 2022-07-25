package main

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"github.com/sirupsen/logrus"
)

var (
	apiClient = client.NewAPIClient(
		settings.NewAuthenticationSettings(
			"",
			"",
		),
		settings.NewHttpSettings(
			"http://localhost:8080/api",
		))
	//taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
	metadataClient   = client.MetadataResourceApiService{APIClient: apiClient}
)

func main() {
	//taskRunner.StartWorker("task1", Task1, 1, time.Millisecond*100)
	//taskRunner.StartWorker("task2", Task2, 1, time.Millisecond*100)

	logrus.Info("Started Workers")

	//taskRunner.WaitWorkers()

	conductorWorkflow := workflow.NewConductorWorkflow(workflowExecutor).
		Name("my_first_workflow").
		Version(1).
		OwnerEmail("developers@orkes.io")

	conductorWorkflow.Add(workflow.NewSimpleTask("simple_task_2", "simple_task_1")).Add(workflow.NewSimpleTask("simple_task_1", "simple_task_2"))

	//Register the workflow with server
	conductorWorkflow.Register(true)

	//Input can be either a map or a struct that is serializable to a JSON map
	workflowInput := map[string]interface{}{}

	workflowId, err := workflowExecutor.StartWorkflow(&model.StartWorkflowRequest{
		Name:  conductorWorkflow.GetName(),
		Input: workflowInput,
	})

	logrus.Info(workflowId)
	logrus.Info(err)
}

//NameAndCity struct that represents the input to the workflow
type NameAndCity struct {
	Name string
	City string
}

func GetTaskDefinitions() []model.TaskDef {
	taskDefs := []model.TaskDef{
		{Name: "task1", TimeoutSeconds: 60},
		{Name: "task2", TimeoutSeconds: 60},
	}
	return taskDefs
}

//NewSimpleWorkflow Create a simple 2-step workflow and register it with the server
/*func NewSimpleWorkflow(executor *executor.WorkflowExecutor) *ConductorWorkflow {

	wf := NewConductorWorkflow(executor).
		Name("simple_workflow").
		Version(1).
		Description("Simple Two Step Workflow").
		TimeoutPolicy(TimeOutWorkflow, 600)

	//Task1
	task1 := NewSimpleTask("task1", "task1").
		Input("name", "${workflow.input.Name}")

	//Task 2
	task2 := NewSimpleTask("task2", "task2").
		Input("city", "${workflow.input.City}")

	//Add two simple tasks
	wf.
		Add(task1).
		Add(task2)

	//Add the output of the workflow from the two tasks
	wf.OutputParameters(map[string]interface{}{
		"Greetings": task1.OutputRef("greetings"),
		"ZipCode":   task2.OutputRef("zip"),
	})

	return wf
}*/

//Task1 worker for Task1
func Task1(task *model.Task) (interface{}, error) {

	//To fail the task send an error
	//return nil, errors.New("bad idea")

	return map[string]interface{}{
		"greetings": "Hello, " + fmt.Sprintf("%v", task.InputData["name"]),
	}, nil
}

//Task2 worker for Task2
func Task2(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"zip": "10121",
	}, nil
}
