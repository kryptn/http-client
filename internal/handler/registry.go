package handler

import (
	"errors"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/kryptn/http-client/language"
)

var blockHandlers = map[string]BlockHandler{}
var functionHandlers = map[string]FunctionHandler{}

func emptyBlockHandler(block language.Block) Result {
	return nil
}

func emptyFunctionHandler(prev Result, fn language.FunctionInvocation) Result {
	return nil
}

func normalizeName(name string) string {
	return strings.ToLower(name)
}

func registerBlockHandler(handlerName string, fn BlockHandler) error {
	normalizedName := normalizeName(handlerName)
	_, ok := blockHandlers[normalizedName]
	if ok {
		return errors.New("a block handler for \"%s\" has already been registered")
	}
	blockHandlers[normalizedName] = fn
	return nil
}

func GetBlockHandler(handlerName string) (BlockHandler, error) {
	normalizedName := normalizeName(handlerName)
	handler, ok := blockHandlers[normalizedName]
	spew.Dump(handler)
	spew.Dump(ok)
	if !ok {
		return emptyBlockHandler, errors.New("block handler \"%s\" not registered")
	}
	return handler, nil
}

func registerFunctionHandler(functionName string, fn FunctionHandler) error {
	normalizedName := normalizeName(functionName)

	_, ok := functionHandlers[normalizedName]
	if ok {
		return errors.New("a function handler for \"%s\" has already been registered")
	}
	functionHandlers[normalizedName] = fn
	return nil
}

func GetFunctionHandler(handlerName string) (FunctionHandler, error) {
	normalizedName := normalizeName(handlerName)
	handler, ok := functionHandlers[normalizedName]
	if !ok {
		return emptyFunctionHandler, errors.New("block handler \"%s\" not registered")
	}
	return handler, nil
}


