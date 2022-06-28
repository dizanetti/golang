package main

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

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
func NewSimpleWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {

	wf := workflow.NewConductorWorkflow(executor).
		Name("simple_workflow").
		Version(1).
		Description("Simple Two Step Workflow").
		TimeoutPolicy(workflow.TimeOutWorkflow, 600)

	//Task1
	task1 := workflow.NewSimpleTask("task1", "task1").
		Input("name", "${workflow.input.Name}")

	//Task 2
	task2 := workflow.NewSimpleTask("task2", "task2").
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
}
