package handler

import (
	"github.com/kryptn/http-client/language"
)

type Result interface {
	Value() interface{}
	HasValue() bool
	Error() error
}

type HandlerFunc func() Result
type BlockHandler func(language.Block) Result
type FunctionHandler func(Result, language.FunctionInvocation) Result

type emptyResult bool

func (er emptyResult) Value() interface{} {
	return true
}

func (er emptyResult) HasValue() bool {
	return false
}

func (er emptyResult) Error() error {
	return nil
}

func HandleBlock(block language.Block) Result {

	handler, err := GetBlockHandler(block.BlockType)
	if err != nil {
		panic("this shouldn't have happened")
	}

	return handler(block)
}

func HandleFunction(prev Result, fn language.FunctionInvocation) Result {

	handler, err := GetFunctionHandler(fn.FuncName)
	if err != nil {
		panic("this shouldn't have happened")
	}
	return handler(prev, fn)
}

func HandleFunctionInvocation(inv []language.FunctionInvocation) Result {
	var result Result
	result = emptyResult(true)

	for _, fn := range inv {
		result = HandleFunction(result, fn)
		if result.Error() != nil {
			return result
		}

	}
	
	return result
}
