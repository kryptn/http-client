package plan

import "github.com/kryptn/http-client/language"

type ExecutionPlan struct {
	documentName     string
	documentFilePath string

	rootBlock *language.Block

	functionInvocations []language.FunctionInvocation
}

func NewExecutionPlan(block *language.Block) *ExecutionPlan {

	return &ExecutionPlan{}
}
